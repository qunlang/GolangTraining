package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println(&name) // 0x82023c080

	changeMe(&name)

	fmt.Println(&name) //0x82023c080
	fmt.Println(name)  //Rocky

	fmt.Println("------其他东西-------") //Rocky
	changeMe111(name)

	fmt.Println(&name) //0x82023c080
	fmt.Println(name)  //Rocky
}

func changeMe(z *string) {
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Todd
	*z = "Rocky"
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Rocky
}

func changeMe111(z string) {
	fmt.Println(z) // 0x82023c080
	z = "Rocky11111"
	fmt.Println(z) // 0x82023c080
}
