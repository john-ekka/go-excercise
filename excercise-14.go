package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min int, max int) int {

	return rand.Intn(max-min) + min

}

func main() {

	rand.Seed(time.Now().Unix())
	randomNum := random(1, 7)
	fmt.Printf("current dice num is : %d", randomNum)
}
