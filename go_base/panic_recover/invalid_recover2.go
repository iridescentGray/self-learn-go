package main

import "fmt"

func main() {
	func() {
		defer func() {
			defer func() {
				fmt.Println(recover()) // 2   恐慌恢复
			}()

			defer recover() // 空操作
			panic(2)
		}()
		panic(1)
	}()
}
