package cli

import (
	"fmt"
)

var (
	// Version ...
	Version string = "9.9.9"
	// Revision =$(git rev-parse --short HEAD)
	Revision string = "9999999"
)

// Parse - parse command line
func Parse() {
	fmt.Printf("version: %s (%s)\n", Version, Revision)
}
