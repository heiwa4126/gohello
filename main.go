package main

import (
	"fmt"
)

var version string
var revision string

func main() {
	fmt.Println("Hello world!")
	fmt.Printf("version: %s-%s\n", version, revision)
}
