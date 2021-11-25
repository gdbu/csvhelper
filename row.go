package csvhelper

import (
	"strings"
)

func newRow(row []string) (r Row, err error) {
	if len(row) == 0 {
		// No bytes exist, return EOF
		err = ErrEmptyRow
		return
	}

	// Make row to match the length of the split
	r = make(Row, len(row))
	// Iterate through split values
	for i, part := range row {
		// Set row value at index of string of split value
		r[i] = strings.ReplaceAll(part, "\\,", ",")
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
