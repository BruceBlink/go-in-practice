package main

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

func main() {
	config, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(config.Section("Section").Key("path").String())
	section, err := config.GetSection("Section")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	enabled, err := section.Key("enabled").Bool()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(enabled)
}
