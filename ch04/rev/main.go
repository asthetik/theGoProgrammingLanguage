package main

import "fmt"

func revers(s []int) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		s[i], s[j] = s[j], s[i]
	}
}

func main() {

	s := []int{0, 1, 2, 3, 4, 5}
	revers(s[:2])
	fmt.Println(s)
	revers(s[2:])
	fmt.Println(s)
	revers(s)
	fmt.Println(s)
}
