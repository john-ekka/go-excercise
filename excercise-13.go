package main

import "fmt"

func main() {

	marks := make(map[string]int)

	marks["maths"] = 1

	fmt.Print(marks["maths"])

	days := map[int]string{
		1: "Sunday",
		2: "Monday",
		3: "Tuesday",
	}

	fmt.Print(days)

	//student-1
	// Maths:33
	// English : 45

	//student-2
	// Maths:33
	// English : 45

	report := map[string]map[string]int{
		"RAMESH": {
			"Maths":   33,
			"English": 45,
		},

		"SURESH": {
			"Maths":   93,
			"English": 95,
		},
	}

	fmt.Print(report["RAMESH"]["English"])
}
