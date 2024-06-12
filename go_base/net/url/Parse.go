package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, _ := url.Parse("https://github.com/SmallP?co=1&bb=2#tag1")
	fmt.Println(u)

	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)
}
