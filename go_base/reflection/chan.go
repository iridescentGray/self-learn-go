package main

import (
	"fmt"
	"reflect"
)

func main() {
	a_chan := make(chan string, 2)
	chan_value := reflect.ValueOf(a_chan)
	fmt.Println(chan_value.Len(), chan_value.Cap()) // 0 2

	//send
	chan_value.Send(reflect.ValueOf("C"))
	succeeded := chan_value.TrySend(reflect.ValueOf("Go"))
	fmt.Println(succeeded) // true
	succeeded = chan_value.TrySend(reflect.ValueOf("C++"))
	fmt.Println(succeeded) // false

	//recv
	vs, succeeded := chan_value.TryRecv()
	fmt.Println(vs.String(), succeeded) // C true
	vs, sentBeforeClosed := chan_value.Recv()
	fmt.Println(vs.String(), sentBeforeClosed) // Go true
	vs, succeeded = chan_value.TryRecv()
	fmt.Println(succeeded) // false
}
