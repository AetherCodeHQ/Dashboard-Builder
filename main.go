package main

import (
	"fmt"
	"os"
)

// dashboard_builder - Build data dashboards
func dashboard_builder(path string) {
	fmt.Println("========================================")
	fmt.Println("  Dashboard-Builder")
	fmt.Println("  Build data dashboards")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	dashboard_builder(path)
}
