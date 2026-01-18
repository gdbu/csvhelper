package csvhelper

import "errors"

var (
	// ErrInvalidRow is returned when a Row contains more entries than the associated Header
	ErrInvalidRow = errors.New("invalid row length, cannot contain more fields than header")
	// ErrEmptyRow is returned when an empty row is encountered
	ErrEmptyRow = errors.New("empty row")
)

var comma = []byte(",")
