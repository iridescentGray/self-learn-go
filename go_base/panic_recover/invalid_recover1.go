package main

import "fmt"

/*
以下所有recover()都无效
*/
func main() {
	defer func() {
		defer func() {
			fmt.Println(recover()) // 空操作
		}()
	}()

	defer func() {
		func() {
			fmt.Println(recover()) // 空操作
		}()
	}()
	func() {
		defer func() {
			fmt.Println(recover()) // 空操作
		}()
	}()
	func() {
		defer fmt.Println(recover()) // 空操作
	}()
	func() {
		fmt.Println(recover()) // 空操作
	}()
	fmt.Println(recover())       // 空操作
	defer fmt.Println(recover()) // 空操作
	panic("bye")
}
