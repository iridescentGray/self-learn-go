package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json_parse = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	parse_result := gjson.Parse(json_parse)
	fmt.Println(parse_result.Index)
	fmt.Println(parse_result)                        //{"name":{"first":"Janet","last":"Prichard"},"age":47}
	fmt.Printf("Type of intVar: %T\n", parse_result) // Type of intVar: gjson.Result

	fmt.Println(parse_result.Get("name"))             //{"first":"Janet","last":"Prichard"}
	fmt.Println(parse_result.Get("name").Get("last")) // Prichard

}
