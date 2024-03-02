package main

import "fmt"

func main() {
	// gradesArray
	fmt.Printf("\n")
	gradesArray()

	// gradesArrayLiteral
	fmt.Printf("\n")
	gradesArrayLiteral()

	// studentArray
	fmt.Printf("\n")
	studentArray()

	//nestedArrays
	fmt.Printf("\n")
	nestedArrays()

	// copyingArrays
	fmt.Printf("\n")
	copyingArrays()

	// copyingArraysWithPointers
	fmt.Printf("\n")
	copyingArraysWithPointers()

	// slicingArrays
	fmt.Printf("\n")
	slicingArrays()

	// slicingArrays2
	fmt.Printf("\n")
	slicingArrays2()

	// sliceWithMakeFunction
	fmt.Printf("\n")
	sliceWithMakeFunction()

	// sliceWithAppendFunction
	fmt.Printf("\n")
	sliceWithAppendFunction()

	// concatenatingSlice
	fmt.Printf("\n")
	concatenatingSlice()

	// stackOperation
	fmt.Printf("\n")
	stackOperation()
}

func gradesArray() {
	fmt.Printf("*** gradesArray ***\n")
	grades := [3]int{89, 90, 45}
	fmt.Printf("Grades: %v", grades)
}

func gradesArrayLiteral() {
	fmt.Printf("*** gradesArrayLiteral ***\n")
	grades := [...]int{89, 90, 45, 34, 90}
	fmt.Printf("Grades: %v", grades)
}

func studentArray() {
	fmt.Printf("*** studentArray ***\n")
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "List"
	students[1] = "Dict"
	students[2] = "Tuple"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of students: %v\n", len(students))
}

func nestedArrays() {
	fmt.Printf("*** nestedArrays ***\n")
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix2)
}

func copyingArrays() {
	fmt.Printf("*** copyingArrays ***\n")
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
}

func copyingArraysWithPointers() {
	fmt.Printf("*** copyingArraysWithPointers ***\n")
	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
}

func slicingArrays() {
	fmt.Printf("*** slicingArrays ***\n")
	a := []int{1, 2, 3}
	b := a
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}

func slicingArrays2() {
	fmt.Printf("*** slicingArrays2 ***\n")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	f := a[len(a)-4:]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

func sliceWithMakeFunction() {
	fmt.Printf("*** sliceWithMakeFunction ***\n")
	a := make([]int, 3, 100)

	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}

func sliceWithAppendFunction() {
	fmt.Printf("*** sliceWithAppendFunction ***\n")
	a := []int{}

	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}

func concatenatingSlice() {
	fmt.Printf("*** concatenatingSlice ***\n")
	a := []int{}

	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, []int{2, 3, 4, 5}...)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}

func stackOperation() {
	fmt.Printf("*** stackOperation ***\n")
	a := []int{1, 2, 3, 4, 5}
	b := a[:len(a)-1]

	fmt.Println(a)
	fmt.Println(b)

	// removing a middle element from slice
	c := append(a[:2], a[3:]...)
	fmt.Println(c)
}
