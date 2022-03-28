package main

import "fmt"

var x = 42

func main() {
	fmt.Println("main: ", x)
	foo()
}

func foo() {
	fmt.Println("foo: ", x)
}
