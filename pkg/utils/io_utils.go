package utils

import (
	"fmt"
	"io"
	"os"
)

func InitFileIO(filename string, action func(string) (*os.File, error)) (io.ReadWriter, func() error, error) {
    file, err := action(filename)
    
    if err != nil {
        return nil, nil, fmt.Errorf("cannot open file: %w", err)
    }
    
    return file, file.Close, nil
}

func PrintError(err error) {
    fmt.Fprintln(os.Stderr, err)
}
