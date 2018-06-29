package main

import "fmt"

//	An	IntList	is	a	linked	list	of	integers.
//	A	nil	*IntList	represents	the	empty	list.
type IntList struct {
	Value int
	Tail  *IntList
}

//	Sum	returns	the	sum	of	the	list	elements.
func (list *IntList) Sum() int {
	// need null check, otherwise panic
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	var list = IntList{3, nil}
	fmt.Println(list.Sum())

}
