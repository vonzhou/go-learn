package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.Tick(500 * time.Millisecond)
	for {
		fmt.Println(<-t1)
	}
}
