package csvhelper

import (
	"bytes"
	"testing"
)

func TestEncoder(t *testing.T) {
	var (
		e   *Encoder
		tss []testStruct
		err error
	)

	tss = append(tss, testStruct{
		FirstName: "John",
		LastName:  "Doe",
		Age:       32,
		City:      "Portland",
		State:     "Oregon",
		Notes: `"Favorite foods:
- eggs
- apples
- pears"`,
	})

	tss = append(tss, testStruct{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       30,
		City:      "Portland",
		State:     "Oregon",
		Notes:     "foo bar",
	})

	// Create buffer
	buf := bytes.NewBuffer(nil)

	// Create header
	header := Row{
		"first_name",
		"last_name",
		"age",
		"city",
		"state",
		"notes",
	}

	// Initialize new encoder
	if e, err = NewEncoder(buf, header); err != nil {
		t.Fatal(err)
	}

	// Iterate through test struct slice
	for _, ts := range tss {
		// Encode test struct
		if err = e.Encode(&ts); err != nil {
			t.Fatal(err)
		}
	}

	// Ensure our output matches the desired output
	if buf.String() != testCSV {
		t.Fatalf("invalid output, expected \"%s\" and received \"%s\"", testCSV, buf.String())
	}
}
