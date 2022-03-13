package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{name: "Rem", age: 23}
	fmt.Println(p1)

	p2 := person{name: "Kim", age: 32}
	fmt.Println(p2)
}