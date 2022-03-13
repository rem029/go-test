package main

import "fmt"

func main() {
	fmt.Println("For loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	a := 0
	
	fmt.Println("While loop")
	for a < 10 {
		fmt.Println(a)
		a++
	}

	arr := []string{"rem","kim","chuchay","onyx"}

	for index,value := range arr{
		fmt.Println("index: ", index, "value",value)
	}

	m := make(map[string]string)

	m["rem"] = "chuchay"
	m["kim"] = "onyx"

	for key,value := range m {
		fmt.Println("key: ",key,"value: ",value)
	}
}