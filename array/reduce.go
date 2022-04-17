package array

import (
	"errors"
	"reflect"
)

func Reduce(source, initialValue, reducer interface{}) (interface{}, error) {
	srcV := reflect.ValueOf(source)
	kind := srcV.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, errors.New("Source must be array")
	}

	if reducer == nil {
		return nil, errors.New("Reducer is not null")
	}

	rv := reflect.ValueOf(reducer)
	if rv.Kind() != reflect.Func {
		return nil, errors.New("Reducer must be function")
	}

	accumulator := initialValue
	accV := reflect.ValueOf(accumulator)

	for i := 0; i < srcV.Len(); i++ {
		entry := srcV.Index(i)

		reduceResults := rv.Call([]reflect.Value{
			accV,
			entry,
			reflect.ValueOf(i),
		})

		accV = reduceResults[0]
	}

	return accV.Interface(), nil
}

func Sum(dataList []Product) int {
	result, _ := Reduce(dataList, 0, func(accumulator int, entry Product, idx int) int {
		return accumulator + int(entry.price)
	})

	if res, ok := result.(int); ok {
		return res
	}

	return 0
}
