package cli

import (
	"bashnya-hw4/pkg/uniq"
	"flag"
	"fmt"
	"os"
)

func ParseFlags() (*uniq.Options, []string, error) {
	options := &uniq.Options{}

	flag.BoolVar(&options.Count, "c", false, "count line")
	flag.BoolVar(&options.Duplicate, "d", false, "print duplicate lines")
	flag.BoolVar(&options.Unique, "u", false, "print uniq lines")
	flag.BoolVar(&options.IgnoreCase, "i", false, "ignore case")
	flag.IntVar(&options.SkipFields, "f", 0, "skip fields")
	flag.IntVar(&options.SkipChars, "s", 0, "skip chars")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]\n")
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if err := options.Validate(); err != nil {
		return nil, nil, err
	}

	files := flag.Args()

	return options, files, nil
}
