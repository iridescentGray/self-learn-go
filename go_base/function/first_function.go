package main

import "fmt"

/**
1.func关键字
2.函数名称。函数名称必须是一个标识符。 这里的函数名称是SquareOfSumAndDiff
3.输入参数列表。输入参数声明列表必须用一对小括号括起来。 输入参数声明有时也称为形参声明（对应后面将介绍的函数调用中的实参）
4.输出结果声明列表
	一个函数可以有多个返回值
5.函数体。函数体必须用一对大括号括起来

**/

// 标准定义
func SquaresOfSumAndDiff(a int64, b int64) (s int64, d int64) {
	s = a * a
	d = b * b
	return // <=> return s, d       返回值可省略
}

//返回值类型可省略
func SquaresOfSumAndDiff2(a int64, b int64) (int64, int64) {
	return (a * a), (b * b)
}

//形参类型可省略
func CompareLower4bits(m, n uint32) (r bool) {
	// 下面这两行等价于：return m&0xFF > n&0xff
	r = m&0xF > n&0xf
	return
}

// 返回值括号可省略
func VersionString() string {
	return "go1.0"
}

// 返回值彻底省略
func doNothing(string, int) {
}

func main() {

	x1, y1 := SquaresOfSumAndDiff(3, 6)
	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println("--------------")
	x2, y2 := SquaresOfSumAndDiff2(3, 6)
	fmt.Println(x2)
	fmt.Println(y2)

}
