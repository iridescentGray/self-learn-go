package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type RPC struct {
	Enable bool   `yaml:"enable"`
	Bind   string `yaml:"bind"`
}

type Config struct {
	RPC        RPC               `yaml:"rpc"`
	Debug      bool              `yaml:"debug"`
	Interval   int               `yaml:"interval"`
	OutPutPath string            `yaml:"out_put_path"`
	Cookies    map[string]string `yaml:"cookies"`
	Rooms      []string          `yaml:"rooms"`
	Numbers    []int             `yaml:"numbers"`
}

func main() {
	file, err := os.ReadFile("./config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	config := Config{}
	if err := yaml.Unmarshal(file, &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)

}
