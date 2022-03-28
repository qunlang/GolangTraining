package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	var i = 0
	for i = 0; i < 100; i++ {
		fmt.Print("第 ", i, " 个; \n")
	}
}
