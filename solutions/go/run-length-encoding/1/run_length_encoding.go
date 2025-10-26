package encode

import (
    "strings"
    "strconv"
    "unicode"
)

func RunLengthEncode(input string) string {
    if len(input) == 0 {
        return ""
    }
	sb := strings.Builder{}
	sb.Grow(len(input))

	i, count := 0, 1
	var cur, prev rune
	for i, cur = range input {
		if i == 0 {
			prev = cur
			continue
		}

		if cur == prev {
			count++
		} else {
			writeToStringBuilder(&sb, count, prev)
			prev = cur
			count = 1
		}
	}

	if count > 0 {
		writeToStringBuilder(&sb, count, cur)
	}

	return sb.String()
}

func writeToStringBuilder(sb *strings.Builder, count int, r rune) {
	if count > 1 {
		sb.WriteString(strconv.Itoa(count))
	}
	sb.WriteRune(r)
}

func RunLengthDecode(s string) string {
	result := strings.Builder{} // to store the resulting string
	result.Grow(100)            // expected string lengths in our system
	length := strings.Builder{} // to store the run-length number
	length.Grow(5)              // we don't expect run-lengths to be more than 5 digits

	// loop over string runes and store digits of run-length in 'length'
	// if character is found, decode the run length and append to 'result'
	for _, r := range s {
		if unicode.IsDigit(r) {
			length.WriteRune(r)
		} else {
			var count int
			if length.Len() > 0 {
				count, _ = strconv.Atoi(length.String())
				length.Reset()
			}

			for count > 1 { // this handles the chars with 1 count
				result.WriteRune(r)
				count--
			}
			result.WriteRune(r)
		}
	}

	return result.String()
}