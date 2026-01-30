package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Привет мир")

	fmt.Println(Bank{})
	fmt.Println(Bank{Name: "LunarBank", BinFrom: 123456, BinTo: 1234567890})

	bytes, err := os.ReadFile("banks.txt")
	if err != nil {
		panic(err)
	}
	text := string(bytes)
	fmt.Println(text)
}

type Bank struct {
	Name    string
	BinFrom int
	BinTo   int
}
