package main

import (
	"fmt"
	"reflect"
)

/*
Select
模拟通道相关的select-case语法
*/
func main() {
	c := make(chan int, 1)
	vc := reflect.ValueOf(c)

	//定义select-case
	select_case_branches := []reflect.SelectCase{
		{Dir: reflect.SelectDefault, Chan: reflect.Value{}, Send: reflect.Value{}},
		{Dir: reflect.SelectRecv, Chan: vc, Send: reflect.Value{}},
		{Dir: reflect.SelectSend, Chan: vc, Send: reflect.ValueOf(789)},
	}

	vc.Send(reflect.ValueOf(123))

	selIndex, vRecv, sentBeforeClosed := reflect.Select(select_case_branches)
	fmt.Println(selIndex)         // 1
	fmt.Println(vRecv.Int())      // 123
	fmt.Println(sentBeforeClosed) // true
}
