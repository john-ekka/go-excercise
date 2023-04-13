package main

import "fmt"

func main() {
	var numbers = []int64{1, 2, 3, 4}

	for _, item := range numbers {
		fmt.Print(item, "\n")
	}
}
