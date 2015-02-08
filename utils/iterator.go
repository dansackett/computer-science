// A stateful iterator in Go which minimally mimics a Generator in Python.
// Usage:
//
// x := []int{1, 2, 3, 4}
// it3, _ := NewIterator(x)
// sum := 0
// for v := it3.Next(); v != nil; v = it3.Next() {
//     fmt.Println(v)
// 	   sum += v.(int)
// }
package iterator

import (
	"errors"
	"reflect"
)

type Iterator struct {
	Current int
	Data    []interface{}
}

func NewIterator(d interface{}) (*Iterator, error) {
	t := reflect.TypeOf(d).Kind()
	if t != reflect.Slice {
		return nil, errors.New("Must be of type slice")
	}

	s := reflect.ValueOf(d)
	i := make([]interface{}, s.Len())

	for x := 0; x < s.Len(); x++ {
		i[x] = s.Index(x).Interface()
	}

	return &Iterator{Current: 0, Data: i}, nil
}

func (it *Iterator) Next() interface{} {
	if it.Current >= len(it.Data) {
		return nil
	}
	v := it.Data[it.Current]
	it.Current++
	return v
}
