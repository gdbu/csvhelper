package main

import (
	"bytes"
	"log"
	"strconv"

	"github.com/Hatch1fy/csvhelper"
)

func main() {
	var (
		u   User
		enc *csvhelper.Encoder
		dec *csvhelper.Decoder
		err error
	)

	u.FirstName = "John"
	u.LastName = "Doe"
	u.Age = 32
	u.City = "Portland"
	u.State = "Oregon"

	buf := bytes.NewBuffer(nil)
	header := csvhelper.Row{
		"first_name",
		"last_name",
		"age",
		"city",
		"state",
	}

	if enc, err = csvhelper.NewEncoder(buf, header); err != nil {
		log.Fatal(err)
	}

	if err = enc.Encode(&u); err != nil {
		log.Fatal(err)
	}

	if dec, err = csvhelper.NewDecoder(buf); err != nil {
		log.Fatal(err)
	}

	var nu User
	if err = dec.Decode(&nu); err != nil {
		log.Fatal(err)
	}

	if u != nu {
		log.Fatalf("values do not match, expected %#v and received %#v", u, nu)
	}
}

// User represents a user
type User struct {
	FirstName string
	LastName  string
	Age       int
	City      string
	State     string
}

// MarshalCSV is a marshaling helper function
func (u *User) MarshalCSV(key string) (value string, err error) {
	switch key {
	case "first_name":
		value = u.FirstName
	case "last_name":
		value = u.LastName
	case "age":
		value = strconv.Itoa(u.Age)
	case "city":
		value = u.City
	case "state":
		value = u.State
	}

	return
}

// UnmarshalCSV is a unmarshaling helper function
func (u *User) UnmarshalCSV(key, value string) (err error) {
	switch key {
	case "first_name":
		u.FirstName = value
	case "last_name":
		u.LastName = value
	case "age":
		u.Age, err = strconv.Atoi(value)
	case "city":
		u.City = value
	case "state":
		u.State = value
	}

	return
}
