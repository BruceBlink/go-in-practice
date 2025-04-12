package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var path string
	var enable bool

	path, err = config.Get("path")
	if err != nil {
		fmt.Println()
	}

	enable, err = config.GetBool("enabled")
	if err != nil {
		fmt.Println("`enable` flag not set in conf.yaml", err)
		return
	}
	fmt.Println("path", path)
	fmt.Println("enable", enable)
}
