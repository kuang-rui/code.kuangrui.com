package main

import (
	"fmt"
	"strings"
)

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "lixiang"
	age = 10
	isOk = true
	fmt.Printf("name:%s", name)
	fmt.Printf("age:%d", age)
	fmt.Println(isOk)

	fmt.Printf("%T\n", name)
	fmt.Printf("%d\n", age)
	fmt.Printf("%v\n", isOk)

	address := "广东佛山siudhfaiu"
	ss := name + address

	fmt.Println(strings.Contains(ss, "lixiang"))
	fmt.Println(strings.HasPrefix(ss, "lixiang"))
	fmt.Println(strings.Index(name,"xi"))
	fmt.Println(strings.LastIndex(name,"i"))

	for _,addr:= range address {
		fmt.Printf("%c\n",addr)
	}

	

}
