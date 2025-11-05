package uniq

import (
	"bufio"
	"fmt"
	"io"
)

type LineInfo struct {
	Data  string
	Count int
}

func Process(reader io.Reader, options Options) ([]LineInfo, error) {
	lineMap, lineOrder, err := scanLines(reader, options)
	if err != nil {
		return nil, err
	}

	return filterLines(lineMap, lineOrder, options), nil
}

func scanLines(reader io.Reader, options Options) (map[string]*LineInfo, []string, error) {
	scanner := bufio.NewScanner(reader)
	var lineOrder []string
	lineMap := make(map[string]*LineInfo)

	for scanner.Scan() {
		line := scanner.Text()
		key := processLine(line, options)

		if key != "" {
			if data, exists := lineMap[key]; exists {
				data.Count++
			} else {
				lineMap[key] = &LineInfo{
					Data:  line,
					Count: 1,
				}
				lineOrder = append(lineOrder, key)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading input")
	}

	return lineMap, lineOrder, nil
}

func filterLines(lineMap map[string]*LineInfo, lineOrder []string, options Options) []LineInfo {
	var result []LineInfo

	for _, key := range lineOrder {
		data := lineMap[key]

		if options.Duplicate {
			if data.Count > 1 {
				result = append(result, *data)
			}
		} else if options.Unique {
			if data.Count == 1 {
				result = append(result, *data)
			}
		} else {
			result = append(result, *data)
		}
	}

	return result
}

func Format(lines []LineInfo, opts Options) []string {
	result := make([]string, 0, len(lines))

	for _, info := range lines {
		var formatted string

		if opts.Count {
			formatted = fmt.Sprintf("%d %s", info.Count, info.Data)
		} else {
			formatted = info.Data
		}
		result = append(result, formatted)
	}

	return result
}
