package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json_get = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Get(json_get, "name.last")
	fmt.Println(value)

	fmt.Println("------GetBytes -------")
	json_byte_for_get := []byte(`{"status":0}`)
	status_code := gjson.GetBytes(json_byte_for_get, "status").Int()
	fmt.Println(status_code)
}
