package main

import "fmt"

func main() {
	// sayMessage
	fmt.Printf("\n")
	sayMessage("Hello Go!")

	// sayMessage2
	fmt.Printf("\n")
	sayMessage2("Hello Go!", 3)

	// sayGreetings
	fmt.Printf("\n")
	sayGreetings("Hello", "Azeez")

	// sayGreetingsWithPointers
	fmt.Printf("\n")
	g := "Greetings"
	n := "Kareem"
	sayGreetingsWithPointers(&g, &n)
	fmt.Println(n)

	// sum
	fmt.Printf("\n")
	s := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(s)

	// sumPointer
	fmt.Printf("\n")
	sPointer := sumPointer(1, 2, 3, 4, 5, 6)
	fmt.Println(*sPointer)

	// sum2
	fmt.Printf("\n")
	s2 := sum2(1, 2, 3, 4, 5, 6)
	fmt.Println(s2)

	// tryDivide
	fmt.Printf("\n")
	tryDivide()

	// anonFuncOne
	fmt.Printf("\n")
	anonFuncOne()

	// anonFuncAsVariable1
	fmt.Printf("\n")
	anonFuncAsVariable1()

	// anonFuncAsVariable2
	fmt.Printf("\n")
	anonFuncAsVariable2()

	// anonFuncAsVariable3
	fmt.Printf("\n")
	anonFuncAsVariable3()

	// methodOneGreet
	fmt.Printf("\n")
	var gr = greeter{
		greeting: "Hi",
		name:     "Namer!",
	}
	fmt.Println(gr.methodOneGreet())
	fmt.Println("This is the modified new name: ", gr.name)

	// methodOneGreetWithPointer
	fmt.Printf("\n")
	var grP = greeter{
		greeting: "Hi",
		name:     "Namer!",
	}
	fmt.Println(grP.methodOneGreetWithPointer())
	fmt.Println("This is the modified new name: ", grP.name)
}

func sayMessage(msg string) {
	fmt.Printf("*** sayMessage ***\n")
	fmt.Println(msg)
}

func sayMessage2(msg string, idx int) {
	fmt.Printf("*** sayMessage2 ***\n")
	for i := 0; i < idx; i++ {
		fmt.Println(msg)
		fmt.Println("The value of the index is", i)
	}
}

func sayGreetings(greeting, name string) {
	fmt.Printf("*** sayGreetings ***\n")
	fmt.Println(greeting, name)
}

func sayGreetingsWithPointers(greeting, name *string) {
	fmt.Printf("*** sayGreetingsWithPointers ***\n")
	fmt.Println(*greeting, *name)
	*name = "Ziyad"
	fmt.Println(*name)
}

func sum(values ...int) int {
	fmt.Printf("*** sum ***\n")
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
	return result
}

func sumPointer(values ...int) *int {
	fmt.Printf("*** sumPointer ***\n")
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
	return &result
}

func sum2(values ...int) (result int) {
	fmt.Printf("*** sum2 ***\n")
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
	return
}

func tryDivide() {
	fmt.Printf("*** tryDivide ***\n")
	d, err := divide(5, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
}

func divide(a, b float64) (float64, error) {
	fmt.Printf("*** divide ***\n")
	if b == 0.0 {
		return 0, fmt.Errorf("Cannot divide by Zero!")
	}
	return a / b, nil
}

func anonFuncOne() {
	fmt.Printf("*** anonFuncOne ***\n")
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func anonFuncAsVariable1() {
	fmt.Printf("*** anonFuncAsVariable1 ***\n")
	f := func() {
		fmt.Printf("Anonymous Function One!")
	}
	f()
}

func anonFuncAsVariable2() {
	fmt.Printf("*** anonFuncAsVariable2 ***\n")
	var f func() = func() {
		fmt.Printf("Anonymous Function One!")
	}
	f()
}

func anonFuncAsVariable3() {
	fmt.Printf("*** anonFuncAsVariable3 ***\n")
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by Zero")
		} else {
			return a / b, nil
		}
	}
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) methodOneGreet() (string, string) {
	fmt.Printf("*** methodOneGreet ***\n")
	return g.greeting, g.name
}

func (g *greeter) methodOneGreetWithPointer() (string, string) {
	fmt.Printf("*** methodOneGreetWithPointer ***\n")
	g.name = "Modified New Name"
	return g.greeting, g.name
}
