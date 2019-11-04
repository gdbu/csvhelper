package csvhelper

import "io"

// NewRowSlices will generate row maps from an io.Reader
func NewRowSlices(r io.Reader) (header Row, rows []RowSlice, err error) {
	var dec *Decoder
	// Initilize a new decoder
	if dec, err = NewDecoder(r); err != nil {
		return
	}

	for err == nil {
		// Initialize a rowslice with an initial capacity of 3
		row := make(RowSlice, 0, 3)

		// Decode the row
		if err = dec.Decode(&row); err != nil {
			break
		}

		// Append the decoded row to the slice
		rows = append(rows, row)
	}

	switch err {
	case nil:
	case io.EOF:
		// Error was "end of file", set to nil
		err = nil

	default:
		return
	}

	header = dec.Header()
	return
}

// RowSlice is the csv slice type
// Note: RowSlice maintains original CSV order
type RowSlice []KV

// MarshalCSV is a CSV encoding helper func
func (r RowSlice) MarshalCSV(key string) (value string, err error) {
	for _, kv := range r {
		if kv.Key != key {
			continue
		}

		value = kv.Value
		return
	}

	return
}

// UnmarshalCSV is a CSV decoding helper func
func (r *RowSlice) UnmarshalCSV(key, value string) (err error) {
	rr := *r
	rr = append(rr, KV{key, value})
	*r = rr
	return
}

// KV represents a key/value entry
type KV struct {
	Key   string
	Value string
}
