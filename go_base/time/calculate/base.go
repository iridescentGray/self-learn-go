package main

import (
	"fmt"
	"time"
)

/*
计算时间
*/
func main() {

	//获取某个月的开始与结束
	startTime, _ := time.Parse("2006-01", "2024-04")
	endTime := startTime.AddDate(0, 1, -1)
	fmt.Println(endTime)
}
