package uniq

import (
    "strings"
    "unicode"
)

func SkipFields(line string, numFields int) string {
    if numFields == 0 {
        return line
    }

    fieldCount := 0
    inField := false
    
    for i, r := range line {
        isSpace := unicode.IsSpace(r)
        
        if !isSpace && !inField {
            fieldCount++
            inField = true
            
            if fieldCount > numFields {
                return line[i:]
            }
        }
        
        if isSpace && inField {
            inField = false
        }
    }

    return ""
}

func SkipChars(line string, numChars int) string {
    if numChars == 0 {
        return line
    }
    
    runes := []rune(line)
    
    if len(runes) <= numChars {
        return ""
    }
    
    return string(runes[numChars:])
}

func NormalizeCase(line string, ignoreCase bool) string {
    if ignoreCase {
        return strings.ToLower(line)
    }
    return line
}
