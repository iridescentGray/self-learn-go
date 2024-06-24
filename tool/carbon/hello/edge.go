package main

import (
	"fmt"
	"os"

	"github.com/golang-module/carbon/v2"
)

func main() {
	fmt.Fprintln(os.Stdout, "StartOfCentury", carbon.Now().StartOfCentury().ToDateTimeString())
	fmt.Fprintln(os.Stdout, "StartOfDecade", carbon.Now().StartOfDecade().ToDateTimeString())
	fmt.Fprintln(os.Stdout, "EndOfDecade", carbon.Now().EndOfDecade().ToDateTimeString())
	fmt.Fprintln(os.Stdout, "StartOfYear", carbon.Now().StartOfYear().ToDateTimeString())
	fmt.Fprintln(os.Stdout, "StartOfQuarter", carbon.Now().StartOfQuarter().ToDateTimeString())
	fmt.Fprintln(os.Stdout, "StartOfMonth", carbon.Now().StartOfMonth().ToDateTimeString())
	fmt.Fprintln(os.Stdout, "StartOfWeek", carbon.Now().StartOfWeek().ToDateTimeString())
	fmt.Fprintln(os.Stdout, "StartOfDay", carbon.Now().StartOfDay().ToDateTimeString())
	fmt.Fprintln(os.Stdout, "StartOfMinute", carbon.Now().StartOfMinute().ToDateTimeString())
}
