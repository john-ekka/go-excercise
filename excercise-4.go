package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	/*reader := bufio.NewReader(os.Stdin)
	fmt.Print("City Name:")
	city, _ := reader.ReadString('\n')
	fmt.Print("You are in : " + city)*/

	//Get a number from the console and check if it’s between 1 and 10.
	READER := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Number: ")
	str1, _ := READER.ReadString('\n')

	//remove new line
	str1 = strings.Replace(str1, "\r\n", "", -1)

	//convert to integer
	num, e := strconv.Atoi(str1)

	if e != nil {
		fmt.Print("Error in Converstion\n" + e.Error())
	}

	if num >= 1 && num <= 10 {
		fmt.Print("Number is in between 10")
	} else {
		fmt.Print("Number is grater then 10")
	}

}
