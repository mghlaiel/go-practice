package main

import (
	"fmt"
)

func sum(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mult(x, y int) int {
	return x * y
}

func div(x, y float64) float64 {
	d := x / y
	return d
}

func modulo(x, y int) int {
	return x % y
}

func main() {
	x, y := 9, 4
	fmt.Println("the sum is", sum(x, y))
	fmt.Println("the difference is", sub(x, y))
	fmt.Println("the multiplication is", mult(x, y))
	fmt.Println("the quotient is", div(float64(x), float64(y)))
	fmt.Println("the remainder is", modulo(x, y))
	s := func(a, b int) int { return a + b }(3, 4)
	fmt.Println(s)
}
