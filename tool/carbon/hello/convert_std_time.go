package main

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon/v2"
)

func main() {
	// 将标准 time.Time 转换成 Carbon
	fmt.Println(carbon.CreateFromStdTime(time.Now()))
	// 将 Carbon 转换成标准 time.Time
	fmt.Println(carbon.Now().StdTime())

	fmt.Printf("StdTime()  %T\n", carbon.Now().StdTime()) // Type
}
