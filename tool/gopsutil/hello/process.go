package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/process"
)

func main() {
	p, _ := process.Processes()

	// almost every return value is a struct
	fmt.Println(p)

}
