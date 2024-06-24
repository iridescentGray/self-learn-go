package main

import (
	"fmt"
	"os"

	"github.com/golang-module/carbon/v2"
)

func main() {
	fmt.Fprintln(os.Stdout, "parseYear", carbon.Parse("2020").ToString())

	fmt.Fprintln(os.Stdout, "parseDate", carbon.Parse("20200805").ToString())
	fmt.Fprintln(os.Stdout, "parseDate", carbon.Parse("2020-08-05").ToString())

	fmt.Fprintln(os.Stdout, "parseDateTime", carbon.Parse("2020-8-5 13:14:15").ToString())
	fmt.Fprintln(os.Stdout, "parseDateTime", carbon.Parse("20200805131415").ToString())

	fmt.Fprintln(os.Stdout, "templete", carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString())
	fmt.Fprintln(os.Stdout, "templete", carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString())
	fmt.Fprintln(os.Stdout, "templete", carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString())

	fmt.Fprintln(os.Stdout, "Layout", carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString())
	fmt.Fprintln(os.Stdout, "Layout", carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString())

}
