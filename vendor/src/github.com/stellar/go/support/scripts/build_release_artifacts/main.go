package main

// See README.md for a description of this script

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"time"

	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/support/log"
)

var extractBinName = regexp.MustCompile(`^(?P<bin>[a-z-]+)-(?P<tag>.+)$`)

var builds = []struct {
	OS   string
	Arch string
}{
	{"darwin", "amd64"},
	{"linux", "amd64"},
	{"linux", "arm"},
	{"windows", "amd64"},
}

func main() {
	log.SetLevel(log.InfoLevel)
	bin, version := extractFromTag(os.Getenv("TRAVIS_TAG"))
	pkg := packageName(bin)

	run("rm", "-rf", "dist/*")

	if bin == "" {
		log.Info("could not extract info from TRAVIS_TAG: skipping artifact packaging")
		os.Exit(0)
	}

	for _, cfg := range builds {
		name := fmt.Sprintf("%s-%s-%s-%s", bin, version, cfg.OS, cfg.Arch)
		dest := filepath.Join("dist", name)

		// make destination directories
		run("mkdir", "-p", dest)
		run("cp", "LICENSE-APACHE.txt", dest)
		run("cp", "COPYING", dest)
		run("cp", filepath.Join(pkg, "README.md"), dest)
		run("cp", filepath.Join(pkg, "CHANGELOG.md"), dest)

		// rebuild the binary with the version variable set
		build(
			fmt.Sprintf("github.com/stellar/go/%s", pkg),
			filepath.Join(dest, bin),
			version,
			cfg.OS,
			cfg.Arch,
		)

		packageArchive(dest, cfg.OS)
	}
}

func build(pkg, dest, version, buildOS, buildArch string) {
	buildTime := time.Now().Format(time.RFC3339)

	timeFlag := fmt.Sprintf("-X github.com/stellar/go/support/app.buildTime=%s", buildTime)
	versionFlag := fmt.Sprintf("-X github.com/stellar/go/support/app.version=%s", version)

	if buildOS == "windows" {
		dest = dest + ".exe"
	}

	cmd := exec.Command("go", "build",
		"-o", dest,
		"-ldflags", fmt.Sprintf("%s %s", timeFlag, versionFlag),
		pkg,
	)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("GOOS=%s", buildOS),
		fmt.Sprintf("GOARCH=%s", buildArch),
	)
	log.Infof("building %s", pkg)

	log.Infof("running: %s", strings.Join(cmd.Args, " "))
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// enableCgo replaces any CGO_ENABLED flags in `env` with CGO_ENABLED=1
func enableCgo(env []string) (ret []string) {
	for _, e := range env {
		if !strings.HasPrefix(e, "CGO_ENABLED") {
			ret = append(ret, e)
		}
	}
	ret = append(ret, "CGO_ENABLED=1")
	return
}

// extractFromTag extracts the name of the binary that should be packaged in the
// course of execution this script as well as the version it should be packaged
// as, based on the name of the tag in the TRAVIS_TAG environment variable.
// Tags must be of the form `NAME-vSEMVER`, such as `horizon-v1.0.0` to be
// matched by this function.
//
// In the event that the TRAVIS_TAG is missing or the match fails, an empty
// string will be returned.
func extractFromTag(tag string) (string, string) {
	match := extractBinName.FindStringSubmatch(tag)
	if match == nil {
		return "", ""
	}

	return match[1], match[2]
}

// packageArchive tars or zips `dest`, depending upon the OS, then removes
// `dest`, in preparation of travis uploading all artifacts to github releases.
func packageArchive(dest, buildOS string) {
	release := filepath.Base(dest)
	dir := filepath.Dir(dest)

	if buildOS == "windows" {
		pop := pushdir(dir)
		// zip $RELEASE.zip $RELEASE/*
		run("zip", "-r", release+".zip", release)
		pop()
	} else {
		// tar -czf $dest.tar.gz -C $DIST $RELEASE
		run("tar", "-czf", dest+".tar.gz", "-C", dir, release)
	}

	run("rm", "-rf", dest)
}

// package searches the `tools` and `services` packages of this repo to find
// the source directory.  This is used within the script to find the README and
// other files that should be packaged with the binary.
func packageName(binName string) string {
	targets := []string{
		filepath.Join("services", binName),
		filepath.Join("tools", binName),
	}

	var result string

	// Note: we do not short circuit this search when we find a valid result so
	// that we can panic when multiple results are found.  The children of
	// /services and /tools should not have name overlap.
	for _, t := range targets {
		_, err := os.Stat(t)

		if os.IsNotExist(err) {
			continue
		}

		if err != nil {
			panic(errors.Wrap(err, "stat failed"))
		}

		if result != "" {
			panic("sourceDir() found multiple results!")
		}

		result = t
	}

	return result
}

// pushdir is a utility function to temporarily change directories.  It returns
// a func that can be called to restore the current working directory to the
// state it was in when first calling pushdir.
func pushdir(dir string) func() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(errors.Wrap(err, "getwd failed"))
	}

	err = os.Chdir(dir)
	if err != nil {
		panic(errors.Wrap(err, "chdir failed"))
	}

	return func() {
		err := os.Chdir(cwd)
		if err != nil {
			panic(errors.Wrap(err, "revert dir failed"))
		}
	}
}

// utility command to run the provided command that echoes any output.  A failed
// command will trigger a panic.
func run(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	log.Infof("running: %s %s", name, strings.Join(args, " "))
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}
