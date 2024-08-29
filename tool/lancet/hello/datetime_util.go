package main

import (
	"fmt"
	"time"

	"github.com/duke-git/lancet/v2/datetime"
)

func main() {
	now := time.Now()
	tomorrow := datetime.AddDay(now, 1)
	yesterday := datetime.AddDay(now, -1)

	fmt.Println(tomorrow.Sub(now))  //24h0m0s
	fmt.Println(yesterday.Sub(now)) // -24h0m0s

	fmt.Println(datetime.GetNowDate()) //2024-08-29
}
