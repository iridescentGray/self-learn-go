package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
)

func main() {
	c, _ := cpu.Info()
	fmt.Println(c)

	t, _ := cpu.Times(false)
	fmt.Println(t)
}
