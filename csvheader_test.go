package csvhelper

import (
	"fmt"
	"strconv"
)

const (
	testCSV = `first_name,last_name,age,city,state
John,Doe,32,Portland,Oregon
Jane,Doe,30,Portland,Oregon
`
)

type testStruct struct {
	FirstName string
	LastName  string
	Age       int
	City      string
	State     string
}

// Validate ensures the values of the struct match the provided values
func (ts *testStruct) Validate(firstName, lastName string, age int, city, state string) (err error) {
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
		return fmt.Errorf("invalid name, expected %s and received %s", state, ts.State)
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
	}

	return
}
