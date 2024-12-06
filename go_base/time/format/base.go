package main

import (
	"fmt"
	"time"
)

/*
格式化时间
*/
func main() {
	now := time.Now()

	// 格式化为 "YYYY-MM-DD HH:mm:ss"
	formatted := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formatted)

	// 格式化为 "YYYY/MM/DD"
	dateOnly := now.Format("2006/01/02")
	fmt.Println("Date Only:", dateOnly)

	// 格式化为 "HH:mm:ss"
	timeOnly := now.Format("15:04:05")
	fmt.Println("Time Only:", timeOnly)

}
