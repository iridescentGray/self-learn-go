package main

import (
	"fmt"
	"reflect"
)

/*
 */
type MyT []interface{ m() }

func (MyT) m() {}

func main() {

	//判断实现
	fmt.Println("-----Implements------")
	Myt_type := reflect.TypeOf(MyT{})
	empty_interface_type := reflect.TypeOf(new(interface{}))

	fmt.Println(Myt_type)                    //main.MyT
	fmt.Println(empty_interface_type)        //*interface {}
	fmt.Println(Myt_type.Kind())             //slice
	fmt.Println(empty_interface_type.Kind()) // ptr
	fmt.Println(Myt_type.Elem())             // interface { main.m() }
	fmt.Println(empty_interface_type.Elem()) // interface {}

	fmt.Println(empty_interface_type.Implements(empty_interface_type.Elem()))        // true
	fmt.Println(empty_interface_type.Elem().Implements(empty_interface_type.Elem())) // true
	fmt.Println(Myt_type.Implements(empty_interface_type.Elem()))                    // true
	fmt.Println(empty_interface_type.Implements(Myt_type.Elem()))                    // false

}
