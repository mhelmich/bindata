package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mhelmich/bindata"
)

var (
	flags           flag.FlagSet
	packageNameFlag = flags.String("package", "", "name of the go package in generated code")
	outputFlag      = flags.String("o", "", "path of the output file")
)

func main() {
	err := parseFlags()
	if err != nil {
		_ = fmt.Errorf("error generating bin data file; %s", err.Error())
		os.Exit(1)
	}

	err = bindata.New(
		flags.Args(),
		bindata.PackageName(*packageNameFlag),
		bindata.OutputFile(*outputFlag),
		bindata.Archiver(bindata.Tar),
		bindata.Compressor(bindata.Bz),
	).Archive()
	if err != nil {
		_ = fmt.Errorf("error generating bin data file; %s", err.Error())
		os.Exit(1)
	}
}

func parseFlags() error {
	err := flags.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	if packageNameFlag == nil || *packageNameFlag == "" {
		return fmt.Errorf("package flag not set")
	}

	if outputFlag == nil || *outputFlag == "" {
		return fmt.Errorf("output flag not set")
	}

	return nil
}
