package csvhelper

import "github.com/Hatch1fy/errors"

const (
	// ErrInvalidRow is returned when a Row contains more entries than the associated Header
	ErrInvalidRow = errors.Error("invalid row length, cannot contain more fields than header")
	// ErrEmptyRow is returned when an empty row is encountered
	ErrEmptyRow = errors.Error("empty row")
)

var comma = []byte(",")
