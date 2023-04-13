package main

func main() {

	//Make a program that counts from 1 to 10.
	for i := 1; i <= 10; i++ {

		println(i)

	}

	var fruits = []string{"Mango", "Orange"}
	var country = []string{"USA", "INDIA"}
	//loop over array
	for i := 0; i < len(fruits); i++ {

		for j := 0; j < len(country); j++ {
			println("fruit: " + fruits[i] + " avaliable in :" + country[j])
		}

	}

}
