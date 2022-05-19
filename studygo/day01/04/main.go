package main

import "fmt"

func main() {
	var testArray [3]int
	var numArray = [3]int{1, 2}
	var cityArray = [3]string{"北京", "上海", "深圳"}

	fmt.Println(testArray[0])
	fmt.Println(numArray[2])
	fmt.Println(cityArray[1])

	var testArray1 = [...]int{1, 2}
	fmt.Println(testArray1[0])

	var numArray1 = [...]string{"aaaa", "bbb"}
	fmt.Println(numArray1)

	var testArray2 = [...]int{1: 1, 2: 100, 6: 200}
	fmt.Println(testArray2)

	for _, v := range cityArray {
		fmt.Println(v)
	}

	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}
}
