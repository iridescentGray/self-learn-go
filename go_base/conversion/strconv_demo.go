package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi("-42")
	fmt.Println("Type of i:", reflect.TypeOf(i))

	s := strconv.Itoa(-42)
	fmt.Println("Type of s:", reflect.TypeOf(s))

	a, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println("Type of a:", reflect.TypeOf(a))
}
