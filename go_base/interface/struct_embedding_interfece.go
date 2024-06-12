package main

import "fmt"

type Qest interface {
	About() string
	GetName() string
}

type Person struct {
	Qest
	Name string
	Age  int
}

func (p *Person) GetName() string {

	return p.Name
}

func main() {
	p := Person{Name: "qq"}

	fmt.Println(p)
	fmt.Println(p.GetName())

	p.About() //panic no implement

}
