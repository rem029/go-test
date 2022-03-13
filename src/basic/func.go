package main

import (
	"errors"
	"fmt"
)

func main() {
	test1Result, test1Error := sum(153,256)
	test2Result, test2Error := sum(-5,-1)

	fmt.Println(test1Result,test1Error)
	fmt.Println(test2Result,test2Error)
}

func sum(x int, y int) (int,error) {
	if x <=0 || y <= 0 {
		return 0, errors.New("Received negative values")
	} else {
		return x + y, nil
	}
}