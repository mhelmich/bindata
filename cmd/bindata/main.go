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
		fmt.Printf("error generating bin data file; %s", err.Error())
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
		fmt.Printf("error generating bin data file; %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Successfully generated bindata file at '%s'", *outputFlag)
}

func parseFlags() error {
	err := flags.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	if packageNameFlag == nil || *packageNameFlag == "" {
		*packageNameFlag = "bindata"
	}

	if outputFlag == nil || *outputFlag == "" {
		*outputFlag = "bindata/bindata.go"
	}

	return nil
}
