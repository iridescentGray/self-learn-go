package main

import "fmt"

// 声明了两个单独的具名常量。（是的，
// 非ASCII字符可以用做标识符。）
const π = 3.1416
const Pi = π // 等价于：const Pi = 3.1416

// 声明了一组具名常量。
const (
	No         = !Yes
	Yes        = true
	MaxDegrees = 360
	Unit       = "弧度"
)

func main() {
	const DoublePi, HalfPi, Unit2 = π * 2, π * 0.5, "度"
	fmt.Println(DoublePi)
	fmt.Println(HalfPi)
	fmt.Println(Unit2)
	fmt.Println(Yes)
	fmt.Println(Unit)
	fmt.Println(MaxDegrees)
}
