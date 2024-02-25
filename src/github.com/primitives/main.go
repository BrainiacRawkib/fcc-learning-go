package main

import "fmt"

func main() {
	var n bool = false

	fmt.Printf("%v, %T\n", n, n)

	i := 1 == 1
	m := 1 == 2

	fmt.Println(i)
	fmt.Println(m)

	var f bool
	fmt.Printf("%v, %T", f, f)
	fmt.Printf("\n\n")

	integers()
	fmt.Printf("\n\n")

	arithmeticOperation()
	fmt.Printf("\n\n")

	arithmeticOperationTypeCasting()
	fmt.Printf("\n\n")

	bitwiseOperators()
	fmt.Printf("\n\n")

	bitShifting()
	fmt.Printf("\n\n")

	floatingNumber()
	fmt.Printf("\n\n")

	floating64()
	fmt.Printf("\n\n")

	complexNumbers()
	fmt.Printf("\n\n")

	stringOperations()
	fmt.Printf("\n\n")

	runeOperations()
	fmt.Printf("\n\n")
}

func integers() {
	fmt.Printf("*** integers ***")
	n := 42
	fmt.Printf("%v, %T\n", n, n)

	var ui int64 = 42
	fmt.Printf("%v, %T\n", ui, ui)

	var ui2 int64 = -42
	fmt.Printf("%v, %T\n", ui2, ui2)
}

func arithmeticOperation() {
	fmt.Printf("*** arithmeticOperation ***\n")
	a := 10
	b := 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
}

func arithmeticOperationTypeCasting() {
	fmt.Printf("*** arithmeticOperationTypeCasting ***\n")
	var a int = 10
	var b int8 = 3

	fmt.Println(a + int(b))
}

func bitwiseOperators() {
	fmt.Printf("*** bitwiseOperators ***\n")
	a := 10
	b := 3

	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
}

func bitShifting() {
	fmt.Printf("*** bitShifting ***\n")
	a := 8 // 2 ^ 3

	fmt.Println(a << 3) // 2 ^ 3 * 2 ^ 3 = 2 ^ 6
	fmt.Println(a >> 3) // 2 ^ 3 / 2 ^ 3 = 2 ^ 0
}

func floatingNumber() {
	fmt.Printf("*** floatingNumber ***\n")
	n := 3.14
	n = 13.7e72
	n = 2.1e14

	fmt.Printf("%v, %T\n", n, n)

	var n2 float64 = 3.14
	n2 = 13.7e72
	n2 = 2.1e14
	fmt.Printf("%v, %T", n2, n2)
}

func floating64() {
	fmt.Printf("*** floating64 ***\n")
	a := 10.2
	b := 3.7

	n := a + b
	n2 := a - b
	n3 := a * b
	n4 := a / b

	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", n2, n2)
	fmt.Printf("%v, %T\n", n3, n3)
	fmt.Printf("%v, %T\n", n4, n4)
}

func complexNumbers() {
	fmt.Printf("*** complexNumbers ***\n")
	var a complex64 = 1 + 2i

	fmt.Printf("%v, %T\n", a, a)

	var b complex64 = 2 + 5.2i
	fmt.Printf("%v, %T\n", b, b)

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	fmt.Printf("%v, %T\n", real(a), real(a))
	fmt.Printf("%v, %T\n", imag(a), imag(a))

	var n complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", n, n)

	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))
}

func stringOperations() {
	fmt.Printf("*** stringOperations ***\n")
	// string refers to utf-8 characters

	s := "this is a string!"
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	s2 := "this is the second string"
	fmt.Println(s + s2)

	fmt.Printf("*** string to bytes ***\n")
	b := []byte(s)
	fmt.Printf("%v, %T", b, b)
}

func runeOperations() {
	fmt.Printf("*** runeOperations ***\n")
	// rune refers to utf-32 characters

	r := 'a'
	fmt.Printf("%v, %T\n", r, r)

	var e rune = 'e'
	fmt.Printf("%v, %T\n", e, e)
}
