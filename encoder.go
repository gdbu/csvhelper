package csvhelper

import (
	"io"
	"strings"
)

// NewEncoder will return a new encoder
func NewEncoder(w io.Writer, header Row) (ep *Encoder, err error) {
	var e Encoder
	if _, err = w.Write(header.Bytes()); err != nil {
		return
	}

	e.w = w
	e.header = header
	ep = &e
	return
}

// Encoder manages encoding
type Encoder struct {
	// Writer for CSV output
	w io.Writer
	// CSV header
	header Row
}

// Header will return a copy of the Encoders's header
func (e *Encoder) Header() (header Row) {
	header = make(Row, len(e.header))
	for i, v := range e.header {
		header[i] = v
	}

	return
}

// Encode will encode a row
func (e *Encoder) Encode(enc Encodee) (err error) {
	var r Row
	for _, key := range e.header {
		var val string
		if val, err = enc.MarshalCSV(key); err != nil {
			return
		}

		val = strings.Replace(val, ",", "\\,", -1)

		r = append(r, val)
	}

	_, err = e.w.Write(r.Bytes())
	return
}

// Encodee is an interface used for Encoding
type Encodee interface {
	MarshalCSV(key string) (value string, err error)
}
