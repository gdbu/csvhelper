package csvhelper

import (
	"io"
	"strconv"
)

// NewRowUintMaps will generate row maps from an io.Reader
func NewRowUintMaps(r io.Reader) (rows []RowUintMap, err error) {
	var dec *Decoder
	if dec, err = NewDecoder(r); err != nil {
		return
	}

	for err == nil {
		row := make(RowUintMap)
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

// RowUintMap is the csv map type
type RowUintMap map[string]uint64

// MarshalCSV is a CSV encoding helper func
func (r RowUintMap) MarshalCSV(key string) (value string, err error) {
	value = strconv.FormatUint(r[key], 10)
	return
}

// UnmarshalCSV is a CSV decoding helper func
func (r RowUintMap) UnmarshalCSV(key, value string) (err error) {
	r[key], err = strconv.ParseUint(value, 10, 64)
	return
}
