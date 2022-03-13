package main

import "fmt"

func main() {
	vertices := make(map[string]int)

	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["dodecagon"] = 12

	fmt.Println(vertices)

	delete(vertices,"square")

	fmt.Println(vertices)
}