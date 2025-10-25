package logs

import (
    "strings"
    "unicode/utf8"
)

const recRune rune = '\u2757'
const searchRune rune = '\U0001F50D'
const weatherRune rune = '\u2600'

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		switch r {
		case recRune:
			{
				return "recommendation"
			}
		case searchRune:
			{
				return "search"
			}
		case weatherRune:
			{
				return "weather"
			}
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
