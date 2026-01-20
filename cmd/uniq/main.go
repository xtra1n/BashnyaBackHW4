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

	reader, closeReader, err := setupReader(files)
	if err != nil {
		utils.PrintError(err)
		return
	}
	if closeReader != nil {
		defer handleClose(closeReader, "reader")
	}

	writer, closeWriter, err := setupWriter(files)
	if err != nil {
		utils.PrintError(err)
		return
	}
	if closeWriter != nil {
		defer handleClose(closeWriter, "writer")
	}

	if err := processAndWrite(reader, writer, options); err != nil {
		utils.PrintError(err)
	}
}

func setupReader(files []string) (io.Reader, func() error, error) {
	if len(files) > 0 {
		return utils.InitFileIO(files[0], os.Open)
	}
	return os.Stdin, nil, nil
}

func setupWriter(files []string) (io.Writer, func() error, error) {
	if len(files) > 1 {
		return utils.InitFileIO(files[1], os.Create)
	}
	return os.Stdout, nil, nil
}

func handleClose(closer func() error, name string) {
	if err := closer(); err != nil {
		utils.PrintError(fmt.Errorf("error closing %s: %w", name, err))
	}
}

func processAndWrite(reader io.Reader, writer io.Writer, options *uniq.Options) error {
	lines, err := uniq.Process(reader, *options)
	if err != nil {
		return err
	}

	formatted := uniq.Format(lines, *options)
	w := bufio.NewWriter(writer)

	for _, line := range formatted {
		fmt.Fprintln(w, line)
	}

	return w.Flush()
}
