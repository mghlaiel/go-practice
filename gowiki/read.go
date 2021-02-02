package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.Name() == "TestPage.txt" {
			content, err := ioutil.ReadFile("TestPage.txt")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("The content of this file: %s", content)
		}
	}
}
