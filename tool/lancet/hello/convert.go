package main

import (
	"fmt"

	"github.com/duke-git/lancet/v2/convertor"
)

func main() {

	fmt.Println(convertor.ToInt("123"))         // Output:
	fmt.Println(convertor.ToInt("-123"))        // -123
	fmt.Println(convertor.ToInt(float64(12.3))) // 12
	fmt.Println(convertor.ToInt("abc"))         // 0 strconv.ParseInt: parsing "": invalid syntax
	fmt.Println(convertor.ToInt(true))          // 0

}
