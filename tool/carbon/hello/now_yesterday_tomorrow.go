package main

import (
	"fmt"
	"os"

	"github.com/golang-module/carbon/v2"
)

func main() {
	fmt.Fprintln(os.Stdout, "Now ", carbon.Now()) // Now  2024-06-24 10:47:05
	fmt.Println(carbon.Now().ToDateString())      //2024-06-24

	fmt.Println(carbon.Now().String()) //2024-06-24 10:47:05

	fmt.Println(carbon.Now().ToDateTimeString()) //2024-06-24 10:47:05

	fmt.Println(carbon.Now().Timestamp()) //1719197225

	fmt.Println(carbon.Now().TimestampMilli()) // /1719197333519

	// 昨天此刻
	fmt.Fprintln(os.Stdout, "Yesterday ", carbon.Yesterday()) // Yesterday  2024-06-23 10:48:53
	// 明天此刻
	fmt.Fprintln(os.Stdout, "Tomorrow ", carbon.Tomorrow()) // Tomorrow  2024-06-25 10:48:53

}
