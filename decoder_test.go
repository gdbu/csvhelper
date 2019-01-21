package csvhelper

import (
	"io"
	"strings"
	"testing"
)

func TestDecoder(t *testing.T) {
	var (
		d   *Decoder
		tss []testStruct
		err error
	)

	// Create a reader from our testCSV string constant
	r := strings.NewReader(testCSV)

	// Initialize a new decoder
	if d, err = NewDecoder(r); err != nil {
		t.Fatal(err)
	}

	// Iterate until we're done or we encountered an error
	for {
		var ts testStruct
		// Decode test struct
		if err = d.Decode(&ts); err != nil {
			// Error encountered, break
			break
		}

		// Append test struct to slice
		tss = append(tss, ts)
	}

	if err != nil && err != io.EOF {
		// We encountered an error that was NOT io.EOF, call Fatal
		t.Fatal(err)
	}

	// Ensure length is proper
	if len(tss) != 2 {
		t.Fatalf("invalid number of rows, expected %d and received %d", 2, len(tss))
	}

	// Validate the first row
	if err = tss[0].Validate("John", "Doe", 32, "Portland", "Oregon", `"Favorite foods:
- eggs
- apples
- pears"`); err != nil {
		t.Fatal(err)
	}

	// Validate the second row
	if err = tss[1].Validate("Jane", "Doe", 30, "Portland", "Oregon", "foo bar"); err != nil {
		t.Fatal(err)
	}
}
