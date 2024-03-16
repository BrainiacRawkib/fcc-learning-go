package main

import "fmt"

func main() {
	// pointerOne
	fmt.Printf("\n")
	pointerOne()

	// pointerArithmetic
	fmt.Printf("\n")
	pointerArithmetic()

	// pointerStruct
	fmt.Printf("\n")
	pointerStruct()

	// pointerStruct2
	fmt.Printf("\n")
	pointerStruct2()

	// pointerStruct3
	fmt.Printf("\n")
	pointerStruct3()

	// pointerArray
	fmt.Printf("\n")
	pointerArray()

	// pointerSlice
	fmt.Printf("\n")
	pointerSlice()

	// pointerMap
	fmt.Printf("\n")
	pointerMap()
}

func pointerOne() {
	fmt.Printf("*** pointerOne ***\n")
	var a int = 42
	var b *int = &a
	fmt.Println(&a, b)
	fmt.Println(a, b)
	fmt.Println(a, *b)
	a = 90
	fmt.Println(a, *b)
	*b = 30
	fmt.Println(a, *b)
}

func pointerArithmetic() {
	fmt.Printf("*** pointerArithmetic ***\n")
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Println("%v %p %p \n", a, b, c)
	fmt.Println("%v %p %p \n", a, *b, *c)
}

type myStruct struct {
	foo int
}

func pointerStruct() {
	fmt.Printf("*** pointerStruct ***\n")
	var ms *myStruct
	fmt.Println(ms)
	ms = &myStruct{foo: 40}
	fmt.Println(ms)
}

func pointerStruct2() {
	fmt.Printf("*** pointerStruct2 ***\n")
	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	ms.foo = 43
	fmt.Println(ms)     // &{43}
	fmt.Println(ms.foo) // 43
}

func pointerStruct3() {
	fmt.Printf("*** pointerStruct3 ***\n")
	var ms *myStruct
	ms = new(myStruct)
	fmt.Println(ms)
	(*ms).foo = 43
	fmt.Println(ms)        // &{43}
	fmt.Println((*ms).foo) // 43
}

func pointerArray() {
	fmt.Printf("*** pointerArray ***\n")
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 43
	fmt.Println(a, b)
}

func pointerSlice() {
	fmt.Printf("*** pointerSlice ***\n")
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 43
	fmt.Println(a, b)
}

func pointerMap() {
	fmt.Printf("*** pointerMap ***\n")
	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
}
