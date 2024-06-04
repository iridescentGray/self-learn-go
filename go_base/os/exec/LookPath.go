package main

import (
	"fmt"
	"os/exec"
)

/*
检查'目标' 是否在path中
*/
func main() {
	path, err := exec.LookPath("go")

	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}
	fmt.Println(path)
}
