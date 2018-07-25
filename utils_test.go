package csvhelper

import (
	"testing"
)

func TestSplitOnChar(t *testing.T) {
	str := `John,Doe,"eggs, bacon, and milk"`
	spl := splitOnChar([]byte(str), ',')

	for i, bs := range spl {
		var expected string
		rowStr := string(bs)
		switch i {
		case 0:
			expected = "John"
		case 1:
			expected = "Doe"
		case 2:
			expected = "\"eggs, bacon, and milk\""
		default:
			t.Fatalf("Expected %d entries and received %d", 3, len(spl))
		}

		if rowStr != expected {
			t.Fatalf("invalid value, expected %s and received %s", expected, rowStr)
		}
	}
}
