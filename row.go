package csvhelper

import (
	"bytes"
	"strings"
)

func newRow(bs []byte) (r Row, err error) {
	if len(bs) == 0 {
		// No bytes exist, return EOF
		err = ErrEmptyRow
		return
	}

	var spl [][]byte
	// Split bytes on comma
	if spl = splitOnChar(bs, ','); len(spl) == 0 {
		// No values exist, return EOF
		err = ErrInvalidRow
		return
	}

	// Make row to match the length of the split
	r = make(Row, len(spl))
	// Iterate through split values
	for i, v := range spl {
		// Convert escaped commas to no longer be escaped
		v = bytes.Replace(v, []byte("\\,"), []byte(","), -1)
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
