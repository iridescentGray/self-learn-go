package main

import (
	"fmt"
	"math/rand"
)

func main() {

	switch n := rand.Intn(100); n % 9 {
	case 0:
		fmt.Println(n, "is a multiple of 9.")
		// 程序不会自动下一个
	case 1, 2, 3:
		fmt.Println(n, "mod 9 is 1, 2 or 3.")
	case 4, 5, 6:
		fmt.Println(n, "mod 9 is 4, 5 or 6.")
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")
	}

	// default可以在任意位置
	fmt.Println("-------default-----------")
	switch n := rand.Intn(3); n {
	default:
		fmt.Println("n == 2")
	case 0:
		fmt.Println("n == 0")
	case 1:
		fmt.Println("n == 1")
	}

	fmt.Println("-------fallthrough-----------")
	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)
		fallthrough // 跳到下个代码块
	case 5, 6, 7, 8:
		// 一个新声明的n，它只在当前分支代码快内可见。
		n := 99
		fmt.Println("n =", n) // 99
		fallthrough
	default:
		// 下一行中的n和第一个分支中的n是同一个变量。
		// 它们均为switch表达式"n"。
		fmt.Println("n =", n)
	}

	fmt.Println("-------omission-----------")
	switch { // <=> switch true {
	case true:
		fmt.Println("hello")
	default:
		fmt.Println("bye")
	}

}
