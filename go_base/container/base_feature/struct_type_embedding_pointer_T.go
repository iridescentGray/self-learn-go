package main

import (
	"fmt"
	"reflect"
)

type Person2 struct {
	Name string
	Age  int
}

func (p Person2) PrintName2() {
	fmt.Println("Name:", p.Name)
}
func (p *Person2) SetAge2(age int) {
	p.Age = age
}

type Singer2 struct {
	*Person2 // 通过内嵌Person类型来扩展
}

func main() {
	var gaga = Singer2{Person2: &Person2{"Gaga", 30}}
	gaga.PrintName2() // Name: Gaga
	gaga.Name = "Lady Gaga"
	gaga.SetAge2(31)
	(&gaga).SetAge2(32)
	gaga.PrintName2()     // Name: Lady Gaga
	fmt.Println(gaga.Age) // 32

	t := reflect.TypeOf(Singer2{}) // the Singer type

	fmt.Println(t, "has", t.NumMethod(), "methods:")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
	}

	pt := reflect.TypeOf(&Singer2{}) // the *Singer type
	fmt.Println(pt, "has", pt.NumMethod(), "methods:")
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", pt.Method(i).Name, "\n")
	}
}
