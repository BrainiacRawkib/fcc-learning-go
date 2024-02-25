package main

import "fmt"

var a int = 34

func main() {
	var i float32 = 42
	fmt.Printf("%v, %T", i, i)

	fmt.Printf("\n\n")

	k := 22.
	fmt.Printf("%v, %T", k, k)

	fmt.Printf("\n\n")

	fmt.Printf("%v, %T", a, a)
}
