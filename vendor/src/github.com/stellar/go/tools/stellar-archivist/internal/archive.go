// Copyright 2016 Stellar Development Foundation and contributors. Licensed
// under the Apache License, Version 2.0. See the COPYING file at the root
// of this distribution or at http://www.apache.org/licenses/LICENSE-2.0

package archivist

import (
	"io"
	"io/ioutil"
	"path"
	"encoding/json"
	"regexp"
	"strconv"
	"net/url"
	"errors"
	"log"
	"bytes"
	"sync"
)

const hexPrefixPat = "/[0-9a-f]{2}/[0-9a-f]{2}/[0-9a-f]{2}/"
const rootHASPath = ".well-known/stellar-history.json"

type CommandOptions struct {
	Concurrency int
	Range Range
	DryRun bool
	Force bool
	Verify bool
	Thorough bool
}

type ConnectOptions struct {
	S3Region string
}

type ArchiveBackend interface {
	Exists(path string) bool
	GetFile(path string) (io.ReadCloser, error)
	PutFile(path string, in io.ReadCloser) error
	ListFiles(path string) (chan string, chan error)
	CanListFiles() bool
}


type Archive struct {
	mutex sync.Mutex
	checkpointFiles map[string](map[uint32]bool)
	allBuckets map[Hash]bool
	referencedBuckets map[Hash]bool

	expectLedgerHashes map[uint32]Hash
	actualLedgerHashes map[uint32]Hash
	expectTxSetHashes map[uint32]Hash
	actualTxSetHashes map[uint32]Hash
	expectTxResultSetHashes map[uint32]Hash
	actualTxResultSetHashes map[uint32]Hash

	missingBuckets int
	invalidBuckets int

	invalidLedgers int
	invalidTxSets int
	invalidTxResultSets int

	backend ArchiveBackend
}

func (a *Archive) GetPathHAS(path string) (HistoryArchiveState, error) {
	var has HistoryArchiveState
	rdr, err := a.backend.GetFile(path)
	if err != nil {
		return has, err
	}
	defer rdr.Close()
	dec := json.NewDecoder(rdr)
	err = dec.Decode(&has)
	return has, err
}

func (a *Archive) PutPathHAS(path string, has HistoryArchiveState, opts *CommandOptions) error {
	if a.backend.Exists(path) && !opts.Force {
		log.Printf("skipping existing " + path)
		return nil
	}
	buf, err := json.MarshalIndent(has, "", "    ")
	if err != nil {
		return err
	}
	return a.backend.PutFile(path,
		ioutil.NopCloser(bytes.NewReader(buf)))
}

func (a *Archive) BucketExists(bucket Hash) bool {
	return a.backend.Exists(BucketPath(bucket))
}

func (a *Archive) CategoryCheckpointExists(cat string, chk uint32) bool {
	return a.backend.Exists(CategoryCheckpointPath(cat, chk))
}

func (a *Archive) GetRootHAS() (HistoryArchiveState, error) {
	return a.GetPathHAS(rootHASPath)
}

func (a *Archive) GetCheckpointHAS(chk uint32) (HistoryArchiveState, error) {
	return a.GetPathHAS(CategoryCheckpointPath("history", chk))
}

func (a *Archive) PutCheckpointHAS(chk uint32, has HistoryArchiveState, opts *CommandOptions) error {
	return a.PutPathHAS(CategoryCheckpointPath("history", chk), has, opts)
}

func (a *Archive) PutRootHAS(has HistoryArchiveState, opts *CommandOptions) error {
	force := opts.Force
	opts.Force = true
	e := a.PutPathHAS(rootHASPath, has, opts)
	opts.Force = force
	return e
}

func (a *Archive) ListBucket(dp DirPrefix) (chan string, chan error) {
	return a.backend.ListFiles(path.Join("bucket", dp.Path()))
}

func (a *Archive) ListAllBuckets() (chan string, chan error) {
	return a.backend.ListFiles("bucket")
}

func (a *Archive) ListAllBucketHashes() (chan Hash, chan error) {
	sch, errs := a.backend.ListFiles("bucket")
	ch := make(chan Hash)
	rx := regexp.MustCompile("bucket" + hexPrefixPat + "bucket-([0-9a-f]{64})\\.xdr\\.gz$")
	errs = makeErrorPump(errs)
	go func() {
		for s := range sch {
			m := rx.FindStringSubmatch(s)
			if m != nil {
				ch <- MustDecodeHash(m[1])
			}
		}
		close(ch)
	}()
	return ch, errs
}

func (a *Archive) ListCategoryCheckpoints(cat string, pth string) (chan uint32, chan error) {
	ext := categoryExt(cat)
	rx := regexp.MustCompile(cat + hexPrefixPat + cat +
		"-([0-9a-f]{8})\\." + regexp.QuoteMeta(ext) + "$")
	sch, errs := a.backend.ListFiles(path.Join(cat, pth))
	ch := make(chan uint32)
	errs = makeErrorPump(errs)

	go func() {
		for s := range sch {
			m := rx.FindStringSubmatch(s)
			if m != nil {
				i, e := strconv.ParseUint(m[1], 16, 32)
				if e == nil {
					ch <- uint32(i)
				} else {
					errs <- errors.New("decoding checkpoint number in filename " + s)
				}
			}
		}
		close(ch)
	}()
	return ch, errs
}

func Connect(u string, opts *ConnectOptions) (*Archive, error) {
	arch := Archive{
		checkpointFiles:make(map[string](map[uint32]bool)),
		allBuckets:make(map[Hash]bool),
		referencedBuckets:make(map[Hash]bool),
		expectLedgerHashes:make(map[uint32]Hash),
		actualLedgerHashes:make(map[uint32]Hash),
		expectTxSetHashes:make(map[uint32]Hash),
		actualTxSetHashes:make(map[uint32]Hash),
		expectTxResultSetHashes:make(map[uint32]Hash),
		actualTxResultSetHashes:make(map[uint32]Hash),
	}
	if opts == nil {
		opts = new(ConnectOptions)
	}
	for _, cat := range Categories() {
		arch.checkpointFiles[cat] = make(map[uint32]bool)
	}
	parsed, err := url.Parse(u)
	if err != nil {
		return &arch, err
	}
	pth := parsed.Path
	if parsed.Scheme == "s3" {
		// Inside s3, all paths start _without_ the leading /
		if len(pth) > 0 && pth[0] == '/' {
			pth = pth[1:]
		}
		arch.backend = MakeS3Backend(parsed.Host, pth, opts)
	} else if parsed.Scheme == "file" {
		pth = path.Join(parsed.Host, pth)
		arch.backend = MakeFsBackend(pth, opts)
	} else if parsed.Scheme == "http" {
		arch.backend = MakeHttpBackend(parsed, opts)
	} else if parsed.Scheme == "mock" {
		arch.backend = MakeMockBackend(opts)
	} else {
		err = errors.New("unknown URL scheme: '" + parsed.Scheme + "'")
	}
	return &arch, err
}

func MustConnect(u string, opts *ConnectOptions) *Archive {
	arch, err := Connect(u, opts)
	if err != nil {
		log.Fatal(err)
	}
	return arch
}
