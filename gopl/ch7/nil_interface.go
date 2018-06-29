package main

import "fmt"

import "bytes"

import "io"

const (
	debug = false
)

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}

	f(buf)

	if debug {
		// use buf
		fmt.Println("----")
	}

}

// buf = false
// non-nil interface containing a nil pointer
func f(out io.Writer) {
	fmt.Printf("%T\n", out)
	if out != nil {
		out.Write([]byte("hello f "))
	}
}
