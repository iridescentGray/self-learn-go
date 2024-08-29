package main

import (
	"fmt"

	"github.com/duke-git/lancet/v2/maputil"
)

func main() {
	m := map[int]string{
		1: "a",
		2: "a",
		3: "b",
		4: "c",
		5: "d",
	}

	fmt.Println(maputil.Keys(m))   // [1 2 3 4 5]
	fmt.Println(maputil.Values(m)) // [a a b c d]
}
