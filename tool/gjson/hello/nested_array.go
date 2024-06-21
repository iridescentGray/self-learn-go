package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	json := ` '{
					"programmers": [
						{
						"firstName": "Janet", 
						"lastName": "McLaughlin", 
						}, {
						"firstName": "Elliotte", 
						"lastName": "Hunter", 
						}, {
						"firstName": "Jason", 
						"lastName": "Harold", 
						}
					]
					}'`
	// array_obj := gjson.Parse(json)
	fmt.Println(gjson.Get(json, "programmers.#.lastName"))

}
