package main

import (
	"fmt"
	"reflect"
)

/*
reflect.ConvertibleTo
判断可转换
*/
func main() {
	var d rune = 1
	var e int32 = 1
	var f float64 = 1

	fmt.Println(reflect.TypeOf(d).ConvertibleTo(reflect.TypeOf(e))) // true
	fmt.Println(reflect.TypeOf(e).ConvertibleTo(reflect.TypeOf(f))) // true
}
