package main

import (
	"log"
	"os"

	_ "./matchers" // 匿名导入，只是为了包初始化
	"./search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
