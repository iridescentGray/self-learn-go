package main

import (
	"fmt"
)

type T struct {
	Name string
}

func (t T) PrintName() {
	fmt.Println("Name:", t.Name)
}
func (t *T) SetName(name string) {
	t.Name = name
}

func main() {
	t := T{Name: "John"}

	// 调用值接收者方法
	t.PrintName()
	// 只需要调t.SetName("")  相当于调用指针属主方法(&t).SetName("")
	t.SetName("Jane")
	t.PrintName()
}
