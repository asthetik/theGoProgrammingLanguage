package main

import (
	"fmt"
)

func main() {

	s := "hello，世界"
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
