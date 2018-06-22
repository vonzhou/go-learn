package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func visit(path string, f os.FileInfo, err error) error {
	// fmt.Printf("Visited: %s\n", path)
	// fmt.Println(f.IsDir())
	if strings.HasSuffix(path, ".exe") {
		fmt.Sprintln("remove %v", path)
		os.Remove(path)
	}
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
