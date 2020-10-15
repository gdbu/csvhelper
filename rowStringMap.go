package csvhelper

import (
	"io"
)

// NewRowStringMaps will generate row maps from an io.Reader
func NewRowStringMaps(r io.Reader) (rows []RowStringMap, err error) {
	var dec *Decoder
	if dec, err = NewDecoder(r); err != nil {
		return
	}

	for err == nil {
		row := make(RowStringMap)
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

// RowStringMap is the csv map type
type RowStringMap map[string]string

// MarshalCSV is a CSV encoding helper func
func (r RowStringMap) MarshalCSV(key string) (value string, err error) {
	value = r[key]
	return
}

// UnmarshalCSV is a CSV decoding helper func
func (r RowStringMap) UnmarshalCSV(key, value string) (err error) {

	r[key] = value
	return
}
