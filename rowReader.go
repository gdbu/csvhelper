package csvhelper

import (
	"bufio"
	"fmt"
	"io"
)

func newRowReader(rdr io.Reader) *rowReader {
	var (
		r  rowReader
		ok bool
	)

	if r.rdr, ok = rdr.(io.RuneReader); !ok {
		r.rdr = bufio.NewReader(rdr)
	}

	return &r
}

type rowReader struct {
	char  rune
	row   []string
	buf   []rune
	state uint8

	rdr io.RuneReader
}

func (r *rowReader) readRow() (row []string, err error) {
	for r.char, _, err = r.rdr.ReadRune(); err == nil; r.char, _, err = r.rdr.ReadRune() {
		fmt.Println("Rune", string(r.char))
		var rowEnd bool
		switch r.state {
		case 0:
			rowEnd, err = r.state0()
		case 1:
			err = r.state1()
		}

		if rowEnd {
			row = r.row
			r.row = nil
			break
		}
	}

	fmt.Println("OH", err, r.row)

	if err == io.EOF && len(r.row) > 0 {
		err = nil
	} else if err != nil {
		return
	}

	if len(r.row) > 0 {
		row = r.row
		r.row = nil
	}

	return
}

func (r *rowReader) state0() (rowEnd bool, err error) {
	switch r.char {
	case ',':
		r.row = append(r.row, string(r.buf))
		r.buf = nil
	case '\n':
		r.row = append(r.row, string(r.buf))
		r.buf = nil
		rowEnd = true
	case '"':
		r.buf = append(r.buf, r.char)
		r.state = 1
	}

	return
}

func (r *rowReader) state1() (err error) {
	switch r.char {
	case '"':
		r.buf = append(r.buf, r.char)
		r.state = 0
	}

	return
}

type onRow func([][]rune) error
