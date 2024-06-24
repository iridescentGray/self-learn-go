package main

import (
	"fmt"
	"os"

	"github.com/golang-module/carbon/v2"
)

func main() {
	fmt.Fprintln(os.Stdout, "DiffInYears", carbon.Now().DiffInYears(carbon.Parse("2020-08-05 13:14:15")))
	fmt.Fprintln(os.Stdout, "DiffAbsInYears", carbon.Now().DiffAbsInYears(carbon.Parse("2020-08-05 13:14:15")))
	fmt.Fprintln(os.Stdout, "DiffInMonths", carbon.Now().DiffInMonths(carbon.Parse("2020-08-05 13:14:15")))
	fmt.Fprintln(os.Stdout, "DiffInDays", carbon.Now().DiffInDays(carbon.Parse("2020-08-05 13:14:15")))

}
