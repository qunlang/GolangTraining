package main

import "fmt"

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	fmt.Printf("%T\n", s)
	info(s)

	fmt.Println("--------------------------------------------------")

	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	fmt.Println("--------------------------------------------------")
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums1 := []int{2, 3, 4}
	sum1 := 0
	for num1 := range nums1 {
		sum1 += num1
	}
	fmt.Println("sum1:", sum1)
}
