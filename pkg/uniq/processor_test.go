package uniq

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestProcessLine(t *testing.T) {
    tests := []struct {
        name     string
        line     string
        options  Options
        expected string
    }{
        {
            name:     "no options",
            line:     "Hello World",
            options:  Options{},
            expected: "Hello World",
        },
        {
            name:     "ignore case only",
            line:     "Hello World",
            options:  Options{IgnoreCase: true},
            expected: "hello world",
        },
        {
            name:     "skip fields only",
            line:     "We love mu\tsic.",
            options:  Options{SkipFields: 1},
            expected: "love mu\tsic.",
        },
        {
            name:     "skip chars only",
            line:     "I love music.",
            options:  Options{SkipChars: 2},
            expected: "love music.",
        },
        {
            name:     "skip fields and chars",
            line:     "We love music.",
            options:  Options{SkipFields: 1, SkipChars: 5},
            expected: "music.",
        },
        {
            name:     "all options combined",
            line:     "We LOVE Music.",
            options:  Options{SkipFields: 1, SkipChars: 5, IgnoreCase: true},
            expected: "music.",
        },
        {
            name:     "skip more fields than available",
            line:     "Hello",
            options:  Options{SkipFields: 5},
            expected: "",
        },
        {
            name:     "skip more chars than available",
            line:     "Hi",
            options:  Options{SkipChars: 10},
            expected: "",
        },
        {
            name:     "multiple spaces between fields",
            line:     "one  two   three",
            options:  Options{SkipFields: 1},
            expected: "two   three",
        },
        {
            name:     "mixed tabs and spaces",
            line:     "first\tsecond  third",
            options:  Options{SkipFields: 2},
            expected: "third",
        },
    }

    for _, testCase := range tests {
        t.Run(testCase.name, func(t *testing.T) {
            result := processLine(testCase.line, testCase.options)
            require.Equal(t, testCase.expected, result)
        })
    }
}
