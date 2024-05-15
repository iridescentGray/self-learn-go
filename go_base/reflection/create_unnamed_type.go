package main

import (
	"fmt"
	"reflect"
)

/*
从Go1.22开始
• 反射可以动态创建无名类型(Array、Chan、Slice、Map、Func、)
		○ 无法动态创建一个接口类型Struct
• 反射无法声明一个新的类型
*/

func main() {

	ta := reflect.ArrayOf(5, reflect.TypeOf(123))
	fmt.Println(ta) // [5]int

	tc := reflect.ChanOf(reflect.SendDir, ta)
	fmt.Println(tc) // chan<- [5]int

	tp := reflect.PointerTo(ta)
	fmt.Println(tp) // *[5]int

	ts := reflect.SliceOf(tp)
	fmt.Println(ts) // []*[5]int

	tm := reflect.MapOf(ta, tc)
	fmt.Println(tm) // map[[5]int]chan<- [5]int

	tf := reflect.FuncOf([]reflect.Type{ta},
		[]reflect.Type{tp, tc}, false)
	fmt.Println(tf) // func([5]int) (*[5]int, chan<- [5]int)

	tt := reflect.StructOf([]reflect.StructField{
		{Name: "Age", Type: reflect.TypeOf("abc")},
	})
	fmt.Println(tt)            // struct { Age string }
	fmt.Println(tt.NumField()) // 1

}
