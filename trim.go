// Package trim provides functions to handle indentation and whitespace in multiline strings.
//
// This package is inspired by similar utility functions in the Kotlin standard library.
package trim

import (
	"math"
	"strings"
	"unicode"
)

// IsBlank returns true if the string s contains only runes for which
// unicode.IsSpace(rune) is true.
func IsBlank(s string) bool {
	return indexNotSpace(s) == -1
}

func indexNotSpace(s string) int {
	return strings.IndexFunc(s, isNotSpace)
}

// isNotSpace returns the negation of [unicode.IsSpace].
func isNotSpace(r rune) bool {
	return !unicode.IsSpace(r)
}

// Indent detects and removes a common minimal indent from all input lines.
//
//   - Identifies the minimum indent across non-blank lines
//   - Removes the identified indent from each line
//   - Removes the first and last lines if they are blank
//   - Blank lines do not influence the detected indent level
//   - Supports different line ending characters (\r\n, \n, \r)
//   - Output always uses \n as line separator
//   - [IsBlank] is used to determine if a string is blank.
//
// Inspired by Kotlin's trimIndent function.
// https://kotlinlang.org/api/core/kotlin-stdlib/kotlin.text/trim-indent.html
func Indent(s string) string {
	lines := split(s) // handles CRLF, LF, leading and trailing blank line.
	if len(lines) == 0 {
		return ""
	}
	// Find the common minimum indent using sentinel.
	const sentinel = math.MaxInt
	indent := sentinel
	for _, line := range lines {
		if i := indexNotSpace(line); 0 <= i && i < indent {
			// Note that blank lines do not affect the detected indent level.
			indent = i
		}
	}
	if 0 < indent && indent != sentinel {
		// If the minimum indent is found, remove them.
		for i, line := range lines {
			if len(line) > indent {
				lines[i] = line[indent:]
			} else {
				// Because the blank lines do not affect the detected indent level,
				// short strings (of blank lines) may exist.
				lines[i] = ""
			}
		}
	}
	return strings.Join(lines, "\n")
}

// Margin removes a common prefix from each line and trims blank lines.
//
//   - Removes leading whitespace and a specified marginPrefix from each line
//   - Uses "|" as default margin delimiter if marginPrefix is an empty string.
//   - Removes the first and last lines if they are blank
//   - Supports different line ending characters (\r\n, \n, \r)
//   - Output always uses \n as line separator
//   - [IsBlank] is used to determine if a string is blank.
//
// Inspired by Kotlin's trimMargin function.
// https://kotlinlang.org/api/core/kotlin-stdlib/kotlin.text/trim-margin.html
func Margin(s, marginPrefix string) string {
	if marginPrefix == "" {
		marginPrefix = "|" // set default prefix if empty
	}
	lines := split(s) // handles CRLF, LF, leading and trailing blank line.
	p := len(marginPrefix)
	for i, line := range lines {
		if j := strings.Index(line, marginPrefix); j >= 0 {
			if IsBlank(line[:j]) {
				// From the beginning of the string s up to prefix contains only
				// whitespace characters, remove it including prefix.
				lines[i] = line[j+p:]
			}
		}
	}
	return strings.Join(lines, "\n")
}

// split returns slice of lines that are separated with \r\n (CRLF), \n (LF),
// or \r (CR) characters and removes the first and the last lines if they are
// blank.
func split(s string) []string {
	s = eolRepl.Replace(s)          // Replace all of CRLF and CR to LF.
	lines := strings.Split(s, "\n") // Split s into lines separated by LF.
	if len(lines) > 0 && IsBlank(lines[0]) {
		// If the first line exists and is blank ...
		// remove the first blank line
		lines = lines[1:]
	}
	if last := len(lines) - 1; last >= 0 && IsBlank(lines[last]) {
		// If the last line exists and is blank,
		// remove the last blank line.
		lines = lines[:last]
	}
	return lines
}

var eolRepl = strings.NewReplacer("\r\n", "\n", "\r", "\n")
