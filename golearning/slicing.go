package main

import "fmt"

func main() {
	s := []int{1, 2, 6, 8, 14, 9}
	fmt.Println(s)
	s = s[1:5]
	fmt.Println(s)
	s = s[2:4]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)
}
