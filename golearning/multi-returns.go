package main

import (
	"fmt"
	"strconv"
)

func f(x, y int64) (i, j string) {
	i, j = strconv.FormatInt(x, 10), strconv.FormatInt(y, 10)
	return
}
func main() {

	a, b := f(2, 5)
	fmt.Printf("the value of a is %s of type %T, and the value of b is %s of type %T.", a, a, b, b)
}
