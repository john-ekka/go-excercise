package main

import "fmt"

func total(x int, y int) int {

	return x / y

}

func main() {

	result := total(5, 0)
	fmt.Print(result)
	defer fmt.Println("ENd of story!")

}
