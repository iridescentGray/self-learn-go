package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	// Create a string-based set of required classes.
	required := mapset.NewSet[string]()
	required.Add("cooking")
	required.Add("biology")
	required.Add("biology")
	fmt.Println(required)

	// Create a string-based set of science classes.
	sciences := mapset.NewSet[string]()
	sciences.Add("biology")
	sciences.Add("chemistry")
	fmt.Println(sciences)

	//包含
	fmt.Printf("Contains  biology %t", sciences.Contains("biology"))
	fmt.Println("")

	//并集
	all := required.Union(required).Union(sciences)
	fmt.Printf("Union all %s", all)
	fmt.Println("")

	//差集
	fmt.Printf("Difference %s", all.Difference(sciences))
	fmt.Println("")
	//交集
	fmt.Printf("Intersect %s", all.Intersect(sciences))

}
