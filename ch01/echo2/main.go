package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("len = %d\n", len(os.Args))
	fmt.Printf("i = %d\ts = %s\n", 0, os.Args[0])

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
