// +build OMIT

package main

import "fmt"

type Vertex struct {
	X int
	Y int
	z int
}

func main() {
	v := Vertex{1, 2, 3}
	v.X = 4
	fmt.Println(v.X)

	// 小写的变量也可以访问
	fmt.Println(v.z)
}
