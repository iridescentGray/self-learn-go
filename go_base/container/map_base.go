package main

import "fmt"

func main() {
	a := map[string]int{}
	fmt.Println(a)

	k, ok := a["123"]
	fmt.Println(k)  //0
	fmt.Println(ok) //存在返回true 	 不存在false
}
