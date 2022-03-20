package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Printf("len = %d\n", len(os.Args))
	fmt.Printf("i = %d\ts = %s\n", 0, os.Args[0]);
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		fmt.Printf("i = %d\ts = %s\n", i, s);
		sep = " "
	}
	fmt.Println(s)
}