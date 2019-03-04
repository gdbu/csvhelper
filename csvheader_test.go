package csvhelper

import (
	"fmt"
	"strconv"
)

const (
	// This value is setup to match how certain situations are expressed in CSV files
	// Although this may look ugly, the argument could be made that CSV has ugly ways of handling
	// certain situations.
	testCSV = `first_name,last_name,age,city,state,notes
John,Doe,32,Portland,Oregon,"Favorite foods:
- eggs
- apples
- pears"
Jane,Doe,30,Portland,Oregon,foo bar
"Jo, the Plumber",Dazini,54,Portland,Oregon,
`
)

type testStruct struct {
	FirstName string
	LastName  string
	Age       int
	City      string
	State     string
	Notes     string
}

// Validate ensures the values of the struct match the provided values
func (ts *testStruct) Validate(firstName, lastName string, age int, city, state, notes string) (err error) {
	if ts.FirstName != firstName {
		return fmt.Errorf("invalid name, expected %s and received %s", firstName, ts.FirstName)
	}

	if ts.LastName != lastName {
		return fmt.Errorf("invalid name, expected %s and received %s", lastName, ts.LastName)
	}

	if ts.Age != age {
		return fmt.Errorf("invalid name, expected %d and received %d", age, ts.Age)
	}

	if ts.City != city {
		return fmt.Errorf("invalid name, expected %s and received %s", city, ts.City)
	}

	if ts.State != state {
		return fmt.Errorf("invalid name, expected \"%s\" and received \"%s\"", state, ts.State)
	}

	if ts.Notes != notes {
		return fmt.Errorf("invalid notes, expected \"%s\" and received \"%s\" / %d / %d", notes, ts.Notes, len(notes), len(ts.Notes))
	}

	return
}

func (ts *testStruct) MarshalCSV(key string) (value string, err error) {
	switch key {
	case "first_name":
		value = ts.FirstName
	case "last_name":
		value = ts.LastName
	case "age":
		value = strconv.Itoa(ts.Age)
	case "city":
		value = ts.City
	case "state":
		value = ts.State
	case "notes":
		value = ts.Notes
	}

	return
}

func (ts *testStruct) UnmarshalCSV(key, value string) (err error) {
	switch key {
	case "first_name":
		ts.FirstName = value
	case "last_name":
		ts.LastName = value
	case "age":
		ts.Age, err = strconv.Atoi(value)
	case "city":
		ts.City = value
	case "state":
		ts.State = value
	case "notes":
		ts.Notes = value
	}

	return
}
