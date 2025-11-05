package uniq

import "strings"

func SkipFields(line string, numFields int) string {
	if numFields <= 0 {
		return line
	}

	fields := strings.Fields(line)

	if numFields >= len(fields) {
		return ""
	}

	return strings.Join(fields[numFields:], " ")
}

func SkipChars(line string, numChars int) string {
	if numChars <= 0 {
		return line
	}

	if numChars >= len(line) {
		return ""
	}

	return string(line[numChars:])
}

func NormalizeCase(line string, ignoreCase bool) string {
	if !ignoreCase {
		return line
	}

	return strings.ToLower(line)
}
