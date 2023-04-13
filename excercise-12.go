package main

import "fmt"

func main() {

	type Person struct {
		name string
		job  string
	}

	type Fruit struct {
		name string
	}

	fr := []Fruit{"apple", "mango"}

	var rep Person

	rep.job = "dev"
	rep.name = "john"

	rep.job = "test"
	rep.name = "laxman"

	fmt.Println("name : " + rep.name + ", job : " + rep.job)
	fmt.Println("name : " + rep.name + ", job : " + rep.job)

}
