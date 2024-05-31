package main

import "fmt"

type Slice []bool

func (s Slice) Length() int {
	return len(s)
}

func (s Slice) Modify(i int, x bool) {
	s[i] = x
}

func main() {
	// 不会造成恐慌
	_ = ((Slice)(nil)).Length
	_ = ((Slice)(nil)).Modify
	// 不会造成恐慌
	_ = ((Slice)(nil)).Length()

	var a Slice = nil
	a.Length()
	fmt.Println(a.Length)

	//下面的调用会产生恐慌，来自于这两个方法内对空指针的解引用
	// ((Slice)(nil)).Modify(0, true)
}
