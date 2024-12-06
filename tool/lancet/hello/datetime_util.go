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

	fmt.Println(datetime.GetNowDate()) //2024-12-03

	//TimeToStr
	fmt.Println(datetime.FormatTimeToStr(now, "yyyy-mm")) //2024-12
	fmt.Println(datetime.FormatTimeToStr(now, "yyyy-mm")) //2024-08-29

	//StrToTime
	fmt.Println(datetime.FormatStrToTime("2024-03", "yyyy-mm")) //2024-12

}
