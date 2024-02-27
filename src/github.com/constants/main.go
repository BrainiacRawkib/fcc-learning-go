package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	a2 = iota
	b2
	c2
)

const (
	_ = iota + 5 // ignore first value by assigning to blank identifier
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHQ
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	const myConst int = 43
	const a string = "string"
	const b float32 = 3.14
	const c bool = true

	fmt.Printf("%v, %T\n", myConst, myConst)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)

	// addConstAndVariable
	fmt.Printf("\n")
	addConstAndVariable()

	// constantInferring
	fmt.Printf("\n")
	constantInferring()

	// iotaConstant
	fmt.Printf("\n")
	iotaConstant()

	// animalConstants
	fmt.Printf("\n")
	animalConstants()

	// fileSizes
	fmt.Printf("\n")
	fileSizes()

	// usersAndRoles
	fmt.Printf("\n")
	usersAndRoles()
}

func addConstAndVariable() {
	fmt.Printf("*** addConstAndVariable ***\n")
	const a int = 42
	var b int = 27
	fmt.Printf("%v, %T\n", a+b, a+b)
}

func constantInferring() {
	fmt.Printf("*** constantInferring ***\n")
	const a = 42
	var b int = 27
	fmt.Printf("%v, %T\n", a+b, a+b)
}

func iotaConstant() {
	fmt.Printf("*** iotaConstants ***\n")
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)

	fmt.Printf("%v, %T\n", a2, a2)
	fmt.Printf("%v, %T\n", b2, b2)
	fmt.Printf("%v, %T\n", c2, c2)
}

func animalConstants() {
	fmt.Printf("*** animalConstants ***\n")
	var specialistType int
	var truthy = specialistType == dogSpecialist
	fmt.Printf("%v, %T\n", truthy, truthy)
	fmt.Printf("dogSpecialist ===> %v\n", dogSpecialist)
}

func fileSizes() {
	fmt.Printf("*** fileSizes ***\n")
	fileSize := 4000000000.
	fmt.Printf("fileSizes ===> %.2fGB\n", fileSize/GB)
}

func usersAndRoles() {
	fmt.Printf("*** usersAndRoles ***\n")
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHQ&roles == isHQ)
}
