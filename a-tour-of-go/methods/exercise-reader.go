package main

import (
	"fmt"
)

type MyReader struct{}

func (r MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 65
	}
	return len(bytes), nil
}

func main() {
	t := MyReader{}
	b := make([]byte, 8)
	n, err := t.Read(b)
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println(b)
}
