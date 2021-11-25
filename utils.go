package csvhelper

import (
	"bytes"
	"strings"
)

func splitOnChar(bs []byte, b byte) (spl [][]byte) {
	var (
		index       int
		escapeState bool
		quoteState  bool
	)

	for i, char := range bs {
		if escapeState {
			// We are currently in an escaped state for this character
			// Escaped state only lasts for one character, set back to false
			escapeState = false
			// This character was escaped, continue
			continue
		}

		switch char {
		case '"':
			// We encounted a double quote, inverse the quoted state
			quoteState = !quoteState
		case '\\':
			// We encounted a backslash, set the escape state to true
			escapeState = true
		case b:
			if quoteState {
				// We cannot split on during an active quote state, continue
				continue
			}

			// Append the part to the split slice
			spl = append(spl, bs[index:i])

			// Update the index
			index = i + 1
		}
	}

	if index < len(bs) {
		spl = append(spl, bs[index:])
	}

	return
}

func trimNewlineSuffix(data []byte) (out []byte) {
	if len(data) == 0 {
		return
	}

	if data[len(data)-1] != '\r' {
		return data
	}

	return data[:len(data)-1]
}

func getCharCount(bs []byte, b byte) (count int) {
	var i int
	for {
		if i = bytes.IndexByte(bs, b); i == -1 {
			return
		}

		count++

		if bs = bs[i+1:]; len(bs) == 0 {
			return
		}
	}
}

func isQuoted(bs []byte) (quoted bool) {
	if getCharCount(bs, '"')%2 != 0 {
		// We have an odd number of double quotes, inverse isQuoted state
		return true
	}

	return
}

func needsEscape(str string) bool {
	switch {
	case strings.Contains(str, ","):
	case strings.Contains(str, "\n"):
	case strings.Contains(str, "\""):

	default:
		return false
	}

	return true
}

func isEscaped(str string) bool {
	if str[0] != '"' {
		return false
	}

	if str[len(str)-1] != '"' {
		return false
	}

	return true
}

func escapeString(str string) string {
	if !needsEscape(str) {
		return str
	}

	str = strings.ReplaceAll(str, "\"", "\\\"")
	return "\"" + str + "\""
}

func unescapeString(str string) string {
	if len(str) == 0 {
		return ""
	}

	str = strings.ReplaceAll(str, "\\\"", "\"")

	if !isEscaped(str) {
		return str
	}

	if len(str) < 2 {
		return ""
	}

	return str[1 : len(str)-1]
}
