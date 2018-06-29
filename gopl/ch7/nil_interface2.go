package main

import "fmt"

import "bytes"

import "io"

const (
	debug = false
)

func main() {
	var buf bytes.Buffer
	if debug {

	}

	f(&buf)

	if debug {
		// use buf
		fmt.Println("----")
	}

}

func f(out io.Writer) {
	fmt.Printf("%T\n", out)
	if out != nil {
		out.Write([]byte("hello f "))
	}
}
