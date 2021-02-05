package prime

import (
	"fmt"
)

func prime(numbers []int) {
	for _, x := range numbers {
		for j := 2; j < x; j++ {
			if x%j == 0 {
				fmt.Printf("The number %v is not prime, since it's also divisible by %v\n", x, j)
				break
			} else if j == x-1 {
				fmt.Printf("The number %v is prime!\n", x)
			}
		}
	}
	return
}
