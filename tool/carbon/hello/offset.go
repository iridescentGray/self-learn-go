package main

import (
	"fmt"
	"os"

	"github.com/golang-module/carbon/v2"
)

func main() {
	fmt.Fprintln(os.Stdout, "AddDecades", carbon.Now().AddDecades(3).ToDateTimeString())
	fmt.Fprintln(os.Stdout, "AddYear", carbon.Now().AddYear().ToDateTimeString())
	fmt.Fprintln(os.Stdout, "SubYear", carbon.Now().SubYear().ToDateTimeString())
	fmt.Fprintln(os.Stdout, "AddYears", carbon.Now().AddYears(3).ToDateTimeString())
	fmt.Fprintln(os.Stdout, "AddQuarters", carbon.Now().AddQuarters(3).ToDateTimeString())
	fmt.Fprintln(os.Stdout, "AddMonths", carbon.Now().AddMonths(3).ToDateTimeString())

}
