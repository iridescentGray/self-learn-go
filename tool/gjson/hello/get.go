package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json_get = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Get(json_get, "name.last")
	fmt.Println(value) //Prichard
}
