package main

import "fmt"

func main() {

	defer fmt.Println("3")
	defer fmt.Println("1")
	fmt.Println("0")
	fmt.Println("2")

}
