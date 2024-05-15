package main

import (
	"fmt"
	"reflect"
)

/*
reflect.TypeOf
从一个任何非接口类型的值创建一个reflect.Type值，此reflect.Type值表示着此非接口值的类型

reflect.Type是接口类型
它的方法能够查看到reflect.Type值所表示的Go类型的各种信息

# Elem 获取Interface/Pointer的基类型
*/
func main() {
	type A = [16]int16
	var a <-chan map[A][]byte
	fmt.Println(reflect.TypeOf(a).Kind()) // map[[16]int16][]uint8
	fmt.Println(reflect.TypeOf(a).Elem()) // chan

	fmt.Println(reflect.TypeOf(a).ChanDir())            // <-chan
	fmt.Println(reflect.TypeOf(a).Elem().Kind())        // map
	fmt.Println(reflect.TypeOf(a).Elem().Key())         // [16]int16
	fmt.Println(reflect.TypeOf(a).Elem().Key().Kind())  // array
	fmt.Println(reflect.TypeOf(a).Elem().Elem().Kind()) //slice

	// 判断可比较
	fmt.Println("-----Comparable------")
	type C = map[int]int
	var b int = 1
	c := C{}
	fmt.Println(reflect.TypeOf(b).Comparable()) // true
	fmt.Println(reflect.TypeOf(c).Comparable()) // false  map不可比较

	//判断可转换
	fmt.Println("-----ConvertibleTo------")
	var d rune = 1
	var e int32 = 1
	var f float64 = 1

	fmt.Println(reflect.TypeOf(d).ConvertibleTo(reflect.TypeOf(e))) // true
	fmt.Println(reflect.TypeOf(e).ConvertibleTo(reflect.TypeOf(f))) // true
}
