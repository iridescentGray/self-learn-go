package main

import (
	"fmt"
	"reflect"
)

/*
MakeFunc
将函数绑定到函数值上
*/

// 给切片的每个int +1
func addSlice(args []reflect.Value) []reflect.Value {
	inSlice, n := args[0], args[0].Len()
	for i := 0; i < n; i++ {
		element := inSlice.Index(i).Int()
		add_element := reflect.ValueOf(element + 1)
		inSlice.Index(i).Set(add_element)
	}
	return args
}

func Bind(p interface{}, f func([]reflect.Value) []reflect.Value) {
	// invert代表着一个函数值
	invert := reflect.ValueOf(p).Elem()
	invert_type := invert.Type()
	fmt.Println(invert_type)
	invert.Set(reflect.MakeFunc(invert_type, f))
}

func main() {
	var invertInts func([]int64) []int64
	Bind(&invertInts, addSlice)

	fmt.Println(invertInts([]int64{2, 3, 5})) // [5 3 2]

}
