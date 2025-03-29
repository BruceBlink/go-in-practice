package main

import (
	"flag" // 引入标准库的 flag包
	"fmt"
)

// 从flag创建一个新的变量
var name = flag.String("name", "World", "A name to say hello to.")

// 直接定义一个变量用于存储flagd的值
var inSpanish bool

func init() {
	//给 inSpanish变量赋值
	flag.BoolVar(&inSpanish, "spanish", false, "Use Spanish language")
	flag.BoolVar(&inSpanish, "s", false, "Use Spanish language.")
	flag.Parse() // parse flag to inSpanish
}

func main() {
	if inSpanish {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s\n", *name)
	}
}
