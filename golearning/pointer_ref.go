package main

import (
	"fmt"
)

type class struct {
	name     string
	students int
}

func newStudents(c *class, n int) *class {
	fmt.Println(c)
	c.students = c.students + n
	return c
}
func main() {
	c := &class{"first grade", 20}
	fmt.Println(*c)
	fmt.Printf("%v", newStudents(c, 5))
	fmt.Printf("%v", c)
}
