package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("for_read")

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	fmt.Println(err)

	fmt.Println(viper.Get("apiVersion"))
	fmt.Println(viper.Get("metadata"))

}
