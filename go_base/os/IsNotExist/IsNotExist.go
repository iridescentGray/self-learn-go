package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("aaa.go")
	fmt.Println(os.IsNotExist(err)) // true

	fmt.Println(err == os.ErrNotExist) // false  错误表现，不推荐使用
}
