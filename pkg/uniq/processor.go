package uniq

func processLine(line string, options Options) string {
	result := line

	result = SkipFields(result, options.SkipFields)
	result = SkipChars(result, options.SkipChars)
	result = NormalizeCase(result, options.IgnoreCase)

	return result
}
