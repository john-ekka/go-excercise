package main

import "fmt"

func main() {

	x, y := 11, 22

	a, b, c := 5, 7, "abc"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(x)
	{
		fmt.Println(x + y)
		x := 10
		fmt.Println(x)
	}
	fmt.Println(x)

	str1, str2 := "Orange", 5
	fmt.Printf("%s %d\n", str1, str2)

	apple, orange := 1, 2.0

	fmt.Printf("total:%d , %.2f", apple, orange)

	//Calculate the year given the date of birth and age

}
