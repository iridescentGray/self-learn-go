package main

import (
	"fmt"
	"math"
)

func main() {
	var a = math.Sqrt(-1.0)
	fmt.Println(a)      // NaN
	fmt.Println(a == a) // false

	var x = 0.0
	var y = 1.0 / x
	var z = 2.0 * y
	fmt.Println(y, z, y == z) // +Inf +Inf true
}
