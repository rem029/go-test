package main

import "fmt"

func main() {
	a := 3
	b := 5
	sum := a+b

	fmt.Println(sum)

	if sum < 7 {
		fmt.Println("more than 7")
	
	}else if sum == 9 {
		fmt.Println("at 8")
	}else {
		fmt.Println("criteria failed")
	}
		
}