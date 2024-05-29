package main

// 下面这行是为了防止f函数的调用被内联。
//go:noinline
func f(i int) byte {
	var a [1 << 20]byte // 使栈增长
	return a[i]
}

/*
同一个变量,地址会改变
*/
func main() {
	var x int
	println(&x) //0xc000057f38
	f(100)
	println(&x) //0xc00027ff38
}
