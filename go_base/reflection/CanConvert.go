package main

import (
	"fmt"
	"reflect"
)

/*
Value.CanConvert(T Type)
检查一个转换是否会成功（即不会产生恐慌）
*/
func main() {
	s := reflect.ValueOf([]int{1, 2, 3, 4, 5})
	ts := s.Type()
	//Go 1.17,切片可以被转化为一个相同元素类型的数组的指针。 但是如果在这样的一个转换中数组类型的长度过长，将导致恐慌产生
	fmt.Println(ts.ConvertibleTo(reflect.TypeOf(&[5]int{}))) // true
	fmt.Println(ts.ConvertibleTo(reflect.TypeOf(&[6]int{}))) // true
	fmt.Println(s.CanConvert(reflect.TypeOf(&[5]int{})))     // true
	fmt.Println(s.CanConvert(reflect.TypeOf(&[6]int{})))     // false
}
