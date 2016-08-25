// Copyright 2016 Stellar Development Foundation and contributors. Licensed
// under the Apache License, Version 2.0. See the COPYING file at the root
// of this distribution or at http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/spf13/cobra"
	archivist "github.com/stellar/go/tools/stellar-archivist/internal"
)

func status(a string, opts *Options) {
	arch := archivist.MustConnect(a, &opts.ConnectOpts)
	state, e := arch.GetRootHAS()
	if e != nil {
		log.Fatal(e)
	}
	buckets := state.Buckets()
	summ, nz := state.LevelSummary()
	fmt.Printf("\n")
	fmt.Printf("       Archive: %s\n", a)
	fmt.Printf("        Server: %s\n", state.Server)
	fmt.Printf(" CurrentLedger: %d (0x%8.8x)\n", state.CurrentLedger, state.CurrentLedger)
	fmt.Printf("CurrentBuckets: %s (%d nonzero levels)\n", summ, nz)
	fmt.Printf(" Newest bucket: %s\n", buckets[0])
	fmt.Printf("\n")
}

type Options struct {
	Low         int
	High        uint32
	Last        int
	Profile     bool
	CommandOpts archivist.CommandOptions
	ConnectOpts archivist.ConnectOptions
}

func (opts *Options) SetRange(arch *archivist.Archive) {
	if arch != nil && opts.Last != -1 {
		state, e := arch.GetRootHAS()
		if e == nil {
			low := state.CurrentLedger - uint32(opts.Last)
			opts.CommandOpts.Range =
				archivist.MakeRange(low, state.CurrentLedger)
			return
		}
	}
	opts.CommandOpts.Range =
		archivist.MakeRange(uint32(opts.Low),
			uint32(opts.High))

}

func (opts *Options) MaybeProfile() {
	if opts.Profile {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}
}

func scan(a string, opts *Options) {
	arch := archivist.MustConnect(a, &opts.ConnectOpts)
	opts.SetRange(arch)
	e1 := arch.Scan(&opts.CommandOpts)
	e2 := arch.ReportMissing(&opts.CommandOpts)
	e3 := arch.ReportInvalid(&opts.CommandOpts)
	if e1 != nil {
		log.Fatal(e1)
	}
	if e2 != nil {
		log.Fatal(e2)
	}
	if e3 != nil {
		log.Fatal(e3)
	}
}

func mirror(src string, dst string, opts *Options) {
	srcArch := archivist.MustConnect(src, &opts.ConnectOpts)
	dstArch := archivist.MustConnect(dst, &opts.ConnectOpts)
	opts.SetRange(srcArch)
	log.Printf("mirroring %v -> %v\n", src, dst)
	e := archivist.Mirror(srcArch, dstArch, &opts.CommandOpts)
	if e != nil {
		log.Fatal(e)
	}
}

func repair(src string, dst string, opts *Options) {
	srcArch := archivist.MustConnect(src, &opts.ConnectOpts)
	dstArch := archivist.MustConnect(dst, &opts.ConnectOpts)
	opts.SetRange(srcArch)
	log.Printf("repairing %v -> %v\n", src, dst)
	e := archivist.Repair(srcArch, dstArch, &opts.CommandOpts)
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	var opts Options

	rootCmd := &cobra.Command{
		Use:   "stellar-archivist",
		Short: "inspect stellar history archive",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	rootCmd.PersistentFlags().IntVar(
		&opts.Low,
		"low",
		0,
		"first ledger to act on",
	)

	rootCmd.PersistentFlags().Uint32Var(
		&opts.High,
		"high",
		uint32(0xffffffff),
		"last ledger to act on",
	)

	rootCmd.PersistentFlags().IntVar(
		&opts.Last,
		"last",
		-1,
		"number of recent ledgers to act on",
	)

	rootCmd.PersistentFlags().IntVarP(
		&opts.CommandOpts.Concurrency,
		"concurrency",
		"c",
		32,
		"number of files to operate on concurrently",
	)

	rootCmd.PersistentFlags().StringVar(
		&opts.ConnectOpts.S3Region,
		"s3region",
		"us-east-1",
		"S3 region to connect to",
	)

	rootCmd.PersistentFlags().BoolVarP(
		&opts.CommandOpts.DryRun,
		"dryrun",
		"n",
		false,
		"describe file-writes, but do not perform any",
	)

	rootCmd.PersistentFlags().BoolVarP(
		&opts.CommandOpts.Force,
		"force",
		"f",
		false,
		"overwrite existing files",
	)

	rootCmd.PersistentFlags().BoolVar(
		&opts.CommandOpts.Verify,
		"verify",
		false,
		"verify file contents",
	)

	rootCmd.PersistentFlags().BoolVar(
		&opts.CommandOpts.Thorough,
		"thorough",
		false,
		"decode and re-encode all buckets",
	)

	rootCmd.PersistentFlags().BoolVar(
		&opts.Profile,
		"profile",
		false,
		"collect and serve profile locally",
	)

	rootCmd.AddCommand(&cobra.Command{
		Use: "status",
		Run: func(cmd *cobra.Command, args []string) {
			status(firstArg(args), &opts)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use: "scan",
		Run: func(cmd *cobra.Command, args []string) {
			opts.MaybeProfile()
			scan(firstArg(args), &opts)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use: "mirror",
		Run: func(cmd *cobra.Command, args []string) {
			opts.MaybeProfile()
			src, dst := srcDst(args)
			mirror(src, dst, &opts)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use: "repair",
		Run: func(cmd *cobra.Command, args []string) {
			opts.MaybeProfile()
			src, dst := srcDst(args)
			repair(src, dst, &opts)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use: "dumpxdr",
		Run: func(cmd *cobra.Command, args []string) {
			err := archivist.DumpXdrAsJson(args)
			if err != nil {
				log.Fatal(err)
			}
		},
	})

	rootCmd.Execute()
}

func firstArg(args []string) string {
	if len(args) == 0 {
		return ""
	}

	return args[0]
}

func srcDst(args []string) (string, string) {
	if len(args) != 2 {
		log.Fatal("require exactly 2 arguments")
	}

	src := args[0]
	dst := args[1]
	if src == "" || dst == "" {
		log.Fatal("require exactly 2 arguments")
	}

	return src, dst
}
