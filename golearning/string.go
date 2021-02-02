package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("enter your name: \n")
	s, _ := reader.ReadString()
	fmt.Printf("the sentence is:\n%s", s)
	fmt.Printf("The sentence has %d characters.", len(s))
	fmt.Printf("your sentence is composed of the following characters:\n ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("the character number %d is %c\n", i+1, s[i])

	}
	s = strings.Join([]string{"hello, there ", "%s\n"}, s)
	fmt.Printf(s)
}
