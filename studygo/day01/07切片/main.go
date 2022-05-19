package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	s := a[1:3]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
}
