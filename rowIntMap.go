package csvhelper

import (
	"io"
	"strconv"
)

// NewRowIntMaps will generate row maps from an io.Reader
func NewRowIntMaps(r io.Reader) (rows []RowIntMap, err error) {
	var dec *Decoder
	if dec, err = NewDecoder(r); err != nil {
		return
	}

	for err == nil {
		row := make(RowIntMap)
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

// RowIntMap is the csv map type
type RowIntMap map[string]int64

// MarshalCSV is a CSV encoding helper func
func (r RowIntMap) MarshalCSV(key string) (value string, err error) {
	value = strconv.FormatInt(r[key], 10)
	return
}

// UnmarshalCSV is a CSV decoding helper func
func (r RowIntMap) UnmarshalCSV(key, value string) (err error) {
	r[key], err = strconv.ParseInt(value, 10, 64)
	return
}
