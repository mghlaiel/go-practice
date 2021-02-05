package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

type Player struct {
	U      User //or simply omit naming a new variable for the struct User
	GameId int
}

func main() {
	p := Player{}
	p.U.Id = 42         // when ommiting the new variable U, p "inherits" access to Id in the struct User. so the assignement operation would look like this: p.Id = 42
	p.U.Name = "Matt"   // p.Name = "Matt"
	p.U.Location = "LA" // p.Location = "LA"
	p.GameId = 90404
	fmt.Printf("%+v", p)
}
