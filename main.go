package main

import (
	"fmt"
)

var Version string
var Revision string

func main() {
	fmt.Println("Hello world!")
	fmt.Printf("version: %s-%s\n", Version, Revision)
}
