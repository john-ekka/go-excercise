package main

import (
	"fmt"
	"os"
)

func main() {

	_, check := os.Stat("excercise-12.go")

	fmt.Println(check)

	if _, err := os.Stat("excercise-11.go"); err == nil {
		fmt.Println("file exists")
	} else {
		fmt.Println("file do not exists")
	}

}
