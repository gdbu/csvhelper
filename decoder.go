package csvhelper

import (
	"bufio"
	"io"
)

// NewDecoder will return a new decoder
func NewDecoder(r io.Reader) (dp *Decoder, err error) {
	var d Decoder
	// Read first line of CSV
	if d.s = bufio.NewScanner(r); !d.s.Scan() {
		err = io.EOF
		return
	}

	// Attempt to parse header from first line bytes
	if d.header, err = newRow(d.s.Bytes()); err != nil {
		return
	}

	dp = &d
	return
}

// Decoder manages decoding
type Decoder struct {
	// Scanner used to read CSV lines
	s *bufio.Scanner
	// CSV header
	header Row
}

// Header will return a copy of the Decoder's header
func (d *Decoder) Header() (header Row) {
	header = make(Row, len(d.header))
	for i, v := range d.header {
		header[i] = v
	}

	return
}

// Decode will decode a single row
func (d *Decoder) Decode(dec Decodee) (err error) {
	// Scan next line
	if !d.s.Scan() {
		// Our scan was unsuccessful which means we've reached the end of our reader, return EOF
		err = io.EOF
		return
	}

	var r Row
	// Attempt to create a new row from our row bytes
	if r, err = newRow(d.s.Bytes()); err != nil {
		if err == ErrEmptyRow {
			return d.Decode(dec)
		}

		return
	}

	// Ensure row length is not longer than header
	if len(r) > len(d.header) {
		return ErrInvalidRow
	}

	// Iterate through row values
	for i, v := range r {
		// Call Decodee's UnmarshalCSV for row value, passing the header entry as the key
		if err = dec.UnmarshalCSV(d.header[i], v); err != nil {
			// Error encountered, return early
			return
		}
	}

	return
}

// Decodee is an interface used for Decoding
type Decodee interface {
	UnmarshalCSV(key, value string) error
}
