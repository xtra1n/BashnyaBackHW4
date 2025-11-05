package main

import (
	"bashnya-hw4/internal/cli"
	"bashnya-hw4/pkg/uniq"
	"bashnya-hw4/pkg/utils"
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	options, files, err := cli.ParseFlags()

	if err != nil {
		utils.PrintError(err)
		return
	}

	var reader io.Reader
	var closeReader func() error

	if len(files) > 0 {
		reader, closeReader, err = utils.InitFileIO(files[0], os.Open)
		if err != nil {
			utils.PrintError(err)
			return
		}
		defer closeReader()
	} else {
		reader = os.Stdin
	}

	var writer io.Writer
	var closeWriter func() error

	if len(files) > 1 {
		writer, closeWriter, err = utils.InitFileIO(files[1], os.Create)
		if err != nil {
			utils.PrintError(err)
			return
		}
		defer closeWriter()
	} else {
		writer = os.Stdout
	}

	lines, err := uniq.Process(reader, *options)
	if err != nil {
		utils.PrintError(err)
		return
	}

	formatted := uniq.Format(lines, *options)

	w := bufio.NewWriter(writer)

	for _, line := range formatted {
		fmt.Fprintln(w, line)
	}

	w.Flush()
}
