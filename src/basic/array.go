package main

import "fmt"

func main() {
	a:= [5]int{}
	a[2] = 5
	fmt.Println(a)

	a = [5]int{1,2,3,4,5}
	fmt.Println(a)

	b := []int{5,4,3,2,1}
	fmt.Println(b)
	b = append(b, 20)
	fmt.Println(b)
}