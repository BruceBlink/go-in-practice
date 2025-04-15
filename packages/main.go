package main

import (
	"fmt"
	"packages/store"
)

func main() {
	product := store.Product{
		Name:     "Kayak",
		Category: "Watersprrts",
	}
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
}
