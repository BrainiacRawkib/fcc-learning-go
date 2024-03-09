package main

import (
	"fmt"
	"math"
)

func main() {
	// statesMapping
	fmt.Printf("\n")
	statePopulations()

	// numberGuessing
	fmt.Printf("\n")
	numberGuessing()

	// ifElse
	fmt.Printf("\n")
	ifElse()

	// compareFloatingNumbers
	fmt.Printf("\n")
	compareFloatingNumbers()

	// switchCase
	fmt.Printf("\n")
	switchCase()

	// switchCase2
	fmt.Printf("\n")
	switchCase2()

	// switchCaseTagList
	fmt.Printf("\n")
	switchCaseTagList()

	// typeSwitch
	fmt.Printf("\n")
	typeSwitch()
}

func statePopulations() {
	fmt.Printf("*** statePopulations ***\n")
	statePopulations := map[string]int{
		"California":   39000000,
		"Texas":        200000,
		"New York":     300000000,
		"Pennsylvania": 120_000_000,
	}
	if pop, ok := statePopulations["Pennsylvania"]; ok {
		fmt.Println(pop)
	}
}

func numberGuessing() {
	fmt.Printf("*** numberGuessing ***\n")
	number := 50
	guess := 30
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low!")
		}
		if guess > number {
			fmt.Println("Too high!")
		}
		if guess == number {
			fmt.Println("Bull's eye!")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
	fmt.Println(!true)
}

func ifElse() {
	fmt.Printf("*** ifElse ***\n")
	number := 50
	guess := 30
	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 1!")
	} else {
		if guess >= 1 && guess <= 100 {
			if guess < number {
				fmt.Println("Too low!")
			}
			if guess > number {
				fmt.Println("Too high!")
			}
			if guess == number {
				fmt.Println("Bull's eye!")
			}
			fmt.Println(number <= guess, number >= guess, number != guess)
		}
		fmt.Println(!true)
	}
}

func compareFloatingNumbers() {
	fmt.Printf("*** compareFloatingNumbers ***\n")
	num := 0.12345
	if math.Abs(num/math.Pow(math.Sqrt(num), 2)-1) < 0.001 {
		fmt.Println("Same numbers!")
	} else {
		fmt.Println("Different numbers!")
	}
}

func switchCase() {
	fmt.Printf("*** switchCase ***\n")
	switch 0 {
	case 1, 3, 5:
		{
			fmt.Println("First case passed!")
		}
	case 2, 4, 6:
		{
			fmt.Println("Second case passed!")
		}
	default:
		fmt.Println("Default case passed!")
	}
}

func switchCase2() {
	fmt.Printf("*** switchCase2 ***\n")
	switch i := 2 + 3; i {
	case 1, 3, 5:
		{
			fmt.Println("First case passed!")
		}
	case 2, 4, 6:
		{
			fmt.Println("Second case passed!")
		}
	default:
		fmt.Println("Default case passed!")
	}
}

func switchCaseTagList() {
	fmt.Printf("*** switchCaseTagList ***\n")
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Tag List 1!")
		fallthrough
	case i >= 20:
		fmt.Println("Tag List 2!")
	default:
		fmt.Println("Default Tag List!")
	}
}

func typeSwitch() {
	fmt.Printf("*** typeSwitch ***\n")
	var i interface{} = 1.
	switch i.(type) {
	case int:
		fmt.Println("i is an integer!")
	case int8:
		fmt.Println("i is an integer8!")
	case int16:
		fmt.Println("i is an integer16!")
	case int32:
		fmt.Println("i is an integer32!")
	case int64:
		fmt.Println("i is an integer64!")
	case string:
		fmt.Println("i is a string!")
	case float32:
		fmt.Println("i is a float32!")
	case float64:
		fmt.Println("i is a float64!")
	default:
		fmt.Println("type not found!")
	}
}
