package main

import (
	"fmt"
	"unsafe"
)

type T1 struct {
	a int8 //1字节

	// 64位架构上，填充7个字节
	// 32位架构上，填充3个字节

	b int64 // 8
	c int16 //2字节
	// 64位架构上，填充6个字节
	// 32位架构上，填充2个字节
}

type T2 struct {
	a int8 //1

	// 64位架构上，填充1个字节
	// 32位架构上，填充1个字节

	c int16 //2

	// 64位架构上，填充4个字节
	// 32位架构上，无需填充

	b int64 //8
}

/*
内存对齐
T1和T2尺寸需为它们的对齐保证的倍数，即在32位架构上为4n个字节，在64位架构上为8n个字节
*/
func main() {

	// 类型T1的尺寸在64位架构上为24个字节（1+7+8+2+6），
	// 在32位架构上为16个字节（1+3+8+2+2）
	fmt.Println(unsafe.Sizeof(T1{}))

	// 类型T2的尺寸在64位架构上位16个字节（1+1+2+4+8），
	// 在32位架构上为12个字节（1+1+2+8）
	fmt.Println(unsafe.Sizeof(T2{}))

	type T3 struct {
		_ [0]int
		b int64 //8
	}
	type T4 struct {
		b int64 //8
	}
	type T5 struct {
		b int64  //8
		_ [0]int //最后一个字段类型的尺寸为零时会影响此结构体的尺寸
	}
	fmt.Println(unsafe.Sizeof(T3{})) //8
	fmt.Println(unsafe.Sizeof(T4{})) //8

	//为了避免取此最后一个零字段的地址越出此结构体内存块的地址，Go编译器会在结构体最后的零尺寸字段后填充字节
	//若结构体的全部都是零尺寸的，那就不需要再填充
	fmt.Println(unsafe.Sizeof(T5{})) //16
}
