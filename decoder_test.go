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

	testCSV := `Title,SKU Number,Category,Specific Category,Designer,Price,Est. Retail Price,Description,LuxRating,Product Keywords,Collections,Origin,Materials,Descriptive Materials,Size,Descriptive Size,Dimension W,Dimension L,Dimension H,Colors,Descriptive Color,Details Colors,Hardware Colors,Condition,Year,Images,Notes,State
Tom Ford Purple Halter Top ,fordt050,apparel,tops,Tom Ford,195.00,800.00,"90s bombshell here we go! This perfectly purple Tom Ford open-back halter top is the quintessential style of nineties mall rat, and back in style today. Made from 100% silk and new with tags, it's TOTALLY for you.",10,"tom ford,cute 90s style shirt,purple,purple halter,grape colored halter top,open back,new with tags,backless shirt,tom ford shirt,women's,tops",new with tags,italy,silk,100% silk,L,L,,31",,purple,,,,Excellent; New or never worn ,,"back::https://luxraise-dev.sfo3.digitaloceanspaces.com/image_1631248275547148049.jpg,front::https://luxraise-dev.sfo3.digitaloceanspaces.com/image_1631248266694228694.jpg",,draft`

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
	if len(tss) != 3 {
		t.Fatalf("invalid number of rows, expected %d and received %d", 3, len(tss))
	}

	// Validate the first row
	if err = tss[0].Validate("John", "Doe", 32, "Portland", "Oregon", `Favorite foods:
- eggs
- apples
- pears`); err != nil {
		t.Fatal(err)
	}

	// Validate the second row
	if err = tss[1].Validate("Jane", "Doe", 30, "Portland", "Oregon", "foo bar"); err != nil {
		t.Fatal(err)
	}

	// Validate the third row
	if err = tss[2].Validate("Jo, the Plumber", "Dazini", 54, "Portland", "Oregon", ""); err != nil {
		t.Fatal(err)
	}
}
