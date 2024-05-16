package main

import (
	"fmt"
	"reflect"
	"time"
)

/*
Implements
判断零值
*/
func main() {
	vx := reflect.ValueOf(123)
	vy := reflect.ValueOf("abc")
	vz := reflect.ValueOf([]bool{false, true})
	vt := reflect.ValueOf(time.Time{})
	fmt.Println(vx.IsZero()) //false
	fmt.Println(vy.IsZero()) //false
	fmt.Println(vz.IsZero()) //false
	fmt.Println(vt.IsZero()) //true
}
