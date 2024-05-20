package main

import (
	"fmt"
	"reflect"
)

func main() {
	valueOf := reflect.ValueOf
	m := map[string]int{"Unix": 1973, "Windows": 1985}
	v := valueOf(m)
	// 第二个实参为Value零值时,表示删除一个key-value
	v.SetMapIndex(valueOf("Windows"), reflect.Value{})
	//新增
	v.SetMapIndex(valueOf("Linux"), valueOf(1991))
	for i := v.MapRange(); i.Next(); {
		fmt.Println(i.Key(), "\t:", i.Value())
	}
}
