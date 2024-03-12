package main

import "fmt"

func main() {
	// loopingOne
	fmt.Printf("\n")
	loopingOne()

	// loopingTwo
	fmt.Printf("\n")
	loopingTwo()

	// loopingThree
	fmt.Printf("\n")
	loopingThree()

	// loopingFour
	fmt.Printf("\n")
	loopingFour()

	// loopingFive
	fmt.Printf("\n")
	loopingFive()

	// loopingSix
	fmt.Printf("\n")
	loopingSix()

	// nestedLoop
	fmt.Printf("\n")
	nestedLoop()

	// loopingSlice
	fmt.Printf("\n")
	loopingSlice()

	// loopingArray
	fmt.Printf("\n")
	loopingArray()

	// loopingMap
	fmt.Printf("\n")
	loopingMap()

	// loopingMap2
	fmt.Printf("\n")
	loopingMap2()

	// loopingString
	fmt.Printf("\n")
	loopingString()
}

func loopingOne() {
	fmt.Printf("*** loopingOne ***\n")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func loopingTwo() {
	fmt.Printf("*** loopingTwo ***\n")
	for i, j := 0, 1; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
}

func loopingThree() {
	fmt.Printf("*** loopingThree ***\n")
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
}

func loopingFour() {
	fmt.Printf("*** loopingFour ***\n")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func loopingFive() {
	fmt.Printf("*** loopingFive ***\n")
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
}

func loopingSix() {
	fmt.Printf("*** loopingSix ***\n")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func nestedLoop() {
	fmt.Printf("*** nestedLoop ***\n")
Loop:
	for i := 1; i < 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
}

func loopingSlice() {
	fmt.Printf("*** loopingSlice ***\n")
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}

func loopingArray() {
	fmt.Printf("*** loopingArray ***\n")
	s := [4]int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}

func loopingMap() {
	fmt.Printf("*** loopingMap ***\n")
	animals := map[string]int{
		"Chicken": 40,
		"Goats":   20,
		"Rams":    40,
		"Cows":    50,
		"Camels":  60,
	}
	for k, v := range animals {
		fmt.Println(k, v)
	}
}

func loopingMap2() {
	fmt.Printf("*** loopingMap2 ***\n")
	animals := map[string]int{
		"Chicken": 40,
		"Goats":   20,
		"Rams":    40,
		"Cows":    50,
		"Camels":  60,
	}
	for _, v := range animals {
		fmt.Println(v)
	}
}

func loopingString() {
	fmt.Printf("*** loopingString ***\n")
	s := "Hello Go world!"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
