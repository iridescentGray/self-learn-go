package main

import (
	"fmt"
	"time"
)

/*
比较时间
*/
func main() {
	now := time.Now()

	// 当天的 18:00 时间
	targetTime := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, now.Location())
	fmt.Println(targetTime)
	// 判断当前时间是否大于或小于 18:00
	if now.Before(targetTime) {
		fmt.Println("当前时间小于 18:00")
	} else {
		fmt.Println("当前时间大于或等于 18:00")
	}
}
