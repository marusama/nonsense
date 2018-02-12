package main

import (
	"errors"
	"regexp"
	"strings"

	"github.com/soider/d"
)

func main() {
	// 1
	//d.D(BuildResultsFixed())

	// 2
	d.D(BuildQueue())

	// 3
	//d.D(BuildRunner())

	// 4
	//r, e := ParseID("123-456")
	//d.D(r, e)

	// 5

	// 6

	// 7
	//http.Handler

}

// 7
// Write the most basic implementation of an http.Handler named "IntegerHandler" that writes an integer (int) value to
// the http.ResponseWriter which was declared during the initiation of the handler.
type IntegerHandler struct{}

// 6

var ln = []byte("\n")

func WriteULN(w *MyWriter) error {
	w.Write(ln)
	return nil

}

func WriteLN(w *MyWriter) error {

	_, err := w.Write([]byte("\n"))
	return err

}

type MyWriter struct {
	s []byte
}

func (m *MyWriter) Write(b []byte) (int, error) {

	m.s = append(m.s, b...)

	return len(b), nil

}

// 5
type CanPerformAction func(uint8) bool

const (
	FlagEdit uint8 = 1 << iota
	FlagRead
	FlagDelete
)

func CanEdit(x uint8) bool {
	return (x & FlagEdit) == FlagEdit
}

func CanRead(x uint8) bool {
	return (x & FlagRead) == FlagRead
}

func CanDelete(x uint8) bool {
	return (x & FlagDelete) == FlagDelete
}

// 4
type ParserType func(s string) ([]string, error)

func ParseID(s string) ([]string, error) {
	matched, err := regexp.MatchString(`^\d[3-5]-\d[3-5]$`, s)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, errors.New("invalid format")
	}
	return strings.SplitN(s, "-", 2), nil
}

// 3
func myRunner() func(int) int {
	return func(x int) int { return x }
}

func BuildRunner() []int {

	rn := myRunner()
	s := []int{0, 1, 2, 3, 5, 8, 13}
	d := []int{}
	for _, x := range s {
		d = append(d, rn(x))
	}
	return d

}

// 2
type Queue []int

func (q *Queue) RemoveLast() {
	r := (*q)[0:(len(*q) - 1)]
	*q = r
}

func BuildQueue() Queue {

	m := &Queue{0, 1, 2, 3}

	m.RemoveLast()

	return *m

}

// 1
type Result struct {
	Data *Data
}

type Data struct {
	ID int
}

func BuildResultsFixed() []Result {
	var list []Result

	for i := 0; i < 10; i++ {
		data := &Data{
			ID: i,
		}
		list = append(list, Result{
			Data: data,
		})
	}
	return list
}
