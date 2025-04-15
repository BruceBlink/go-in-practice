package main

import (
	"fmt"
	"sort"
)

//"math/rand"

func main() {
	//fmt.Println("Value:", rand.Int())
	//const price float32 = 275.00
	//const tax float32 = 27.50
	//fmt.Println(price + tax)
	/**
	first := 100
	second := &first
	first ++
	fmt.Println("first: ", first)  //   101
	fmt.Println("second: ", *second) // 101
	**/
	/**
	first := 100
	second := &first

	first++
	*second++

	fmt.Println("first: ", first)  //   102
	fmt.Println("second: ", *second) // 102
	var myNewPoint = second
	*myNewPoint ++

	fmt.Println("first: ", first)  //   103
	fmt.Println("second: ", *second) // 103
     **/
	 names := [3]string {"Alice", "Charlie", "Bob"}
	 secondPosition := &names[1]
	 fmt.Println(*secondPosition)
	 sort.Strings(names[:])
	 fmt.Println(*secondPosition)
}
