package csvhelper

import (
	"io"
)

// NewRowMaps will generate row maps from an io.Reader
func NewRowMaps(r io.Reader) (rows []RowMap, err error) {
	var dec *Decoder
	if dec, err = NewDecoder(r); err != nil {
		return
	}

	for err == nil {
		row := make(RowMap)
		if err = dec.Decode(row); err != nil {
			break
		}

		rows = append(rows, row)
	}

	if err == io.EOF {
		err = nil
	}

	return
}

// RowMap is the csv map type
type RowMap map[string]string

// MarshalCSV is a CSV encoding helper func
func (r RowMap) MarshalCSV(key string) (value string, err error) {
	value = r[key]
	return
}

// UnmarshalCSV is a CSV decoding helper func
func (r RowMap) UnmarshalCSV(key, value string) (err error) {

	r[key] = value
	return
}
