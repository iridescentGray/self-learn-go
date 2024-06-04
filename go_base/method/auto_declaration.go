package main

import (
	"fmt"
	"reflect"
)

type S struct {
	Name string
}

func (t S) PrintName() {
	fmt.Println(" Name:", t.Name)
}
func (t *S) SetName(name string) {
	t.Name = name
}

/*
值属主方法(PrintName) 将会自动声明一个指针属主方法
*/
func main() {
	reflect_t := reflect.TypeOf(S{})
	fmt.Println(reflect_t, "has", reflect_t.NumMethod(), "methods:")
	for i := 0; i < reflect_t.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", reflect_t.Method(i).Name, "\n")
	}

	pt := reflect.TypeOf(&S{}) // the *Singer type
	fmt.Println(pt, "has", pt.NumMethod(), "methods:")
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", pt.Method(i).Name, "\n")
	}
}
