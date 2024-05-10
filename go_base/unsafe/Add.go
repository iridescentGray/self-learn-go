package main

import (
	"fmt"
	"unsafe"
)

/*
Go 1.17引入
func Add(ptr Pointer, len IntegerType) Pointer

len 代表偏移距离
在一个（非安全）指针表示的地址上添加一个偏移量，然后返回表示新地址的一个指针
*/
func main() {
	a := [10]int{3: 3, 9: 9}
	fmt.Println(a)

	point_9 := &a[9]
	unsafe_point_9 := unsafe.Pointer(point_9)
	//-6 代表向前偏移6*元素数组长度, 等价于 a[3]
	p3 := (*int)(unsafe.Add(unsafe_point_9, -6*int(unsafe.Sizeof(a[0]))))
	fmt.Println(*p3) // 3

}
