package main

import (
	"fmt"

	"gihub.com/heiwa4126/gohello/calc"
)

var (
	Version  string
	Revision string
)

func main() {
	fmt.Println("Hello world!")
	fmt.Printf("version: %s (%s)\n", Version, Revision)

	a := 1
	b := 2
	fmt.Printf("%d + %d = %d\n", a, b, calc.Add(a, b))
	fmt.Printf("%d + %d = %d\n", a, b, calc.Sub(a, b))
}
