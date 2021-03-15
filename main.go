package main

import (
	"fmt"

	"github.com/heiwa4126/gohello/cli"
	"github.com/heiwa4126/gohello/lib"
)

func sayAnything() {
	fmt.Println("Hello world!")
}

func main() {
	sayAnything()
	cli.Parse()

	a := 1
	b := 2
	fmt.Printf("%d + %d = %d\n", a, b, lib.Add(a, b))
	fmt.Printf("%d + %d = %d\n", a, b, lib.Sub(a, b))

	fmt.Printf("Runinng on %s\n", lib.Msg())
}
