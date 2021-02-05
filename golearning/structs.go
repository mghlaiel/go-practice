package main

import (
	"fmt"
	"time"
)

type event struct {
	Location string
	Date     time.Time
}

func main() {
	e := new(event) // initializes a zeroed struct simlar to e := &event{}
	e.Location = "virtual"
	e.Date = time.Now()
	fmt.Println(*e, &e, e)
}
