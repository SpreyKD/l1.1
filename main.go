package main

import "fmt"

type Human struct {
	Age  int
	Name string
}

type Action struct {
	Human
	column int
}

func people() Human {
	return Human{
		Age:  16,
		Name: "danil",
	}
}

func act() {
	a := Action{
		Human:  people(),
		column: 52,
	}
	fmt.Println(a)
}

func main() {
	act()
}
