package main

import (
	"fmt"
	"os/exec"
)

/*
执行外部命令。这个包提供了丰富的功能，用于执行命令并捕获其输出、传递输入等
*/
func main() {

	convertCmd := exec.Command(
		"git",
	)

	out, _ := convertCmd.Output()
	fmt.Println(string(out))
}
