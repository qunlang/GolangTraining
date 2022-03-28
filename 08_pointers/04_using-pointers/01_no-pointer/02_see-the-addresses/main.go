package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z) // address in func zero
	fmt.Println(&z)        // address in func zero
	z = 0
}

func zero1(z *int) {
	fmt.Printf("%p\n", &z) // address in func zero
	fmt.Println(&z)        // address in func zero
	*z = 0
	fmt.Println(*z) // address in func zero
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println(x) // x is still 5
	zero1(&x)
	fmt.Println(x) // x is still 5
}
