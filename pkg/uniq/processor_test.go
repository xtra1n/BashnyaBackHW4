package uniq

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestProcessLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		options     Options
		expected string
	}{
		{
			name:     "no options",
			line:     "Hello World",
			options:     Options{},
			expected: "Hello World",
		},
		{
			name:     "ignore case only",
			line:     "Hello World",
			options:     Options{IgnoreCase: true},
			expected: "hello world",
		},
		{
			name:     "skip fields only",
			line:     "We love music.",
			options:     Options{SkipFields: 1},
			expected: "love mu	sic.",
		},
		{
			name:     "skip chars only",
			line:     "I love music.",
			options:     Options{SkipChars: 1},
			expected: " love music.",
		},
		{
			name:     "skip fields and chars",
			line:     "We love music.",
			options:     Options{SkipFields: 1, SkipChars: 1},
			expected: "ove music.",
		},
		{
			name:     "all options combined",
			line:     "We LOVE Music.",
			options:     Options{SkipFields: 1, SkipChars: 1, IgnoreCase: true},
			expected: "ove music.",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := processLine(testCase.line, testCase.options)
			require.Equal(t, testCase.expected, result)
		})
	}
}
