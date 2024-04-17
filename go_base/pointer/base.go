package main

import "fmt"

func main() {
	p0 := new(int)   // p0指向一个int类型的零值
	fmt.Println(p0)  // （十六进制形式的地址）
	fmt.Println(*p0) // 0

	x := *p0         // x是p0所引用的值的一个复制。
	p1, p2 := &x, &x // p1和p2中都存储着x的地址。
	// x、*p1和*p2表示着同一个int值。
	fmt.Println(p1 == p2) // true
	fmt.Println(p0 == p1) // false

	p3 := &*p0
	fmt.Println(p3)       // （十六进制形式的地址）
	fmt.Println(p0 == p3) // true

	fmt.Printf("%T, %T \n", *p0, x) // int, int
	fmt.Printf("%T, %T \n", p0, p1) // *int, *int
}
