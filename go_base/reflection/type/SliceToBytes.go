package main

import (
	"bytes"
	"fmt"
	"reflect"
)

type MyByte byte

/*
Go类型系统禁止切片类型[]MyByte的值转换为类型[]byte
但eflect.Value的Bytes方法可以帮我们绕过这个限制
*/
func main() {
	var mybs = []MyByte{'a', 'b', 'c'}
	var bs = reflect.ValueOf(mybs).Bytes()
	fmt.Println(bytes.HasPrefix(bs, []byte{'a', 'b'})) // true

	bs[1], bs[2] = 'r', 't'
	fmt.Printf("%s \n", mybs) // art
}
