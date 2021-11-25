package csvhelper

import (
	"errors"
	"fmt"
	"io"
)

var errBreak = errors.New("break")

// NewDecoder will return a new decoder
func NewDecoder(r io.Reader) (dp *Decoder, err error) {
	var d Decoder
	d.r = newRowReader(r)

	var row []string
	if row, err = d.r.readRow(); err != nil {
		return
	}

	if d.header, err = newRow(row); err != nil {
		return
	}

	fmt.Println("Decoding with a header length of", len(d.header))

	dp = &d
	return
}

// Decoder manages decoding
type Decoder struct {
	// Scanner used to read CSV lines
	r *rowReader
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
	var row []string
	if row, err = d.r.readRow(); err != nil {
		fmt.Println("ERR", err)
		return
	}

	fmt.Println("Row", row)

	var r Row
	// Attempt to create a new row from our row bytes
	if r, err = newRow(row); err != nil {
		if err == ErrEmptyRow {
			return d.Decode(dec)
		}

		return
	}

	fmt.Println("Row len?", len(r), r[0])

	// Ensure row length is not longer than header
	if len(r) > len(d.header) {
		return ErrInvalidRow
	}

	// Iterate through row values
	for i, v := range r {
		v = unescapeString(v)
		// Call Decodee's UnmarshalCSV for row value, passing the header entry as the key
		if err = dec.UnmarshalCSV(d.header[i], v); err != nil {
			// Error encountered, return early
			return
		}
	}

	return
}

// ForEach is helper func for iterating through decoder rows
func (d *Decoder) ForEach(fn func() Decodee) (err error) {
	for err == nil {
		if err = d.Decode(fn()); err != nil {
			break
		}
	}

	if err == io.EOF {
		err = nil
	}

	return
}

// Decodee is an interface used for Decoding
type Decodee interface {
	UnmarshalCSV(key, value string) error
}
