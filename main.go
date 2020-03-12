package main

import (
	"fmt"

	"github.com/heiwa4126/gohello/calc"
)

var (
	// Version ...
	Version string = "v9.9.9"
	// Revision =$(git rev-parse --short HEAD)
	Revision string = "9999999"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Printf("version: %s (%s)\n", Version, Revision)

	a := 1
	b := 2
	fmt.Printf("%d + %d = %d\n", a, b, calc.Add(a, b))
	fmt.Printf("%d + %d = %d\n", a, b, calc.Sub(a, b))
}
