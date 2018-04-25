package csvhelper

import (
	"bytes"
	"io"
	"strings"
)

func newRow(b []byte) (r Row, err error) {
	if len(b) == 0 {
		// No bytes exist, return EOF
		err = io.EOF
		return
	}

	var spl [][]byte
	// Split bytes on comma
	if spl = bytes.Split(b, comma); len(spl) == 0 {
		// No values exist, return EOF
		err = io.EOF
		return
	}

	// Make row to match the length of the split
	r = make(Row, len(spl))
	// Iterate through split values
	for i, v := range spl {
		// Set row value at index of string of split value
		r[i] = string(v)
	}

	return
}

// Row represents a CSV row
type Row []string

// String will return the string representation of a row
func (r Row) String() string {
	return strings.Join(r, ",") + "\n"
}

// Bytes will return the bytes representation of a row
func (r Row) Bytes() []byte {
	return []byte(r.String())
}
