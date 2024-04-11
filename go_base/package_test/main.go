package main

import (
	"fmt"
	myrand_package "math/rand"
)

// <=> import "math/rand"

func init() {
	fmt.Println("hi,", bob)
}

func init() {
	fmt.Println("hello,", smith)
}

func main() {
	fmt.Println("bye")
	fmt.Print(" myrand_package  %s", myrand_package.Uint32())
}

func titledName(who string) string {
	return "Mr. " + who
}

var bob, smith = titledName("Bob"), titledName("Smith")
