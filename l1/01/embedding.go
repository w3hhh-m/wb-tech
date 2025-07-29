package main

import "fmt"

// Given a structure Human (with any set of fields and methods).
// Implement embedding of parent structure Human into the structure Action

type Human struct {
	Name string
	Age  int
}

func (h Human) Greeting() {
	fmt.Printf("Hi, my name is %s. I am %d years old.\n", h.Name, h.Age)
}

type Action struct {
	// embedding struct
	Human
	Name       string
	Difficulty int
}

func (a Action) Start() {
	// Action can access Human fields
	fmt.Printf("Human %s (%d y.o.) : starting '%s' task with %d difficulty level.\n", a.Human.Name, a.Human.Age, a.Name, a.Difficulty)
}

func main() {
	human := Human{
		Name: "w3hhh-m",
		Age:  18,
	}

	action := Action{
		Human:      human,
		Name:       "Practicing Golang",
		Difficulty: 5,
	}

	// Action can access Human methods
	action.Greeting()

	action.Start()
}
