package main

import "fmt"

func main() {
	var array = [...]int{1, 3, 5, 7, 8}
	var sum int = 0
	for i := 0; i < len(array); i++ {
		sum = sum + array[i]
	}
	fmt.Println(sum)


	for i := 0; i < len(array)/2; i++ {
		for j := 0; j <len(array); j++ {
			if array[i]+array[j]==8 {
				fmt.Println(i,j)
			}
		}

	}
}
