package csvhelper

import (
	"io"
	"strconv"
)

// NewRowFloatMaps will generate row maps from an io.Reader
func NewRowFloatMaps(r io.Reader) (rows []RowFloatMap, err error) {
	var dec *Decoder
	if dec, err = NewDecoder(r); err != nil {
		return
	}

	for err == nil {
		row := make(RowFloatMap)
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

// RowFloatMap is the csv map type
type RowFloatMap map[string]float64

// MarshalCSV is a CSV encoding helper func
func (r RowFloatMap) MarshalCSV(key string) (value string, err error) {
	value = strconv.FormatFloat(r[key], 'f', 4, 64)
	return
}

// UnmarshalCSV is a CSV decoding helper func
func (r RowFloatMap) UnmarshalCSV(key, value string) (err error) {
	r[key], err = strconv.ParseFloat(value, 64)
	return
}
