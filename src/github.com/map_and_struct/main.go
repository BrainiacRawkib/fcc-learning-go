package main

import (
	"fmt"
	"reflect"
)

func main() {
	// statesMapping
	fmt.Printf("\n")
	statesMapping()

	// statesMappingWithMakeFunc
	fmt.Printf("\n")
	statesMappingWithMakeFunc()

	// commaOkSyntax
	fmt.Printf("\n")
	commaOkSyntax()

	// structDoctor
	fmt.Printf("\n")
	structDoctor()

	// anonymousStruct
	fmt.Printf("\n")
	anonymousStruct()

	// structComposition
	fmt.Printf("\n")
	structComposition()

	// structTags
	fmt.Printf("\n")
	structTags()
}

func statesMapping() {
	fmt.Printf("*** statesMapping ***\n")
	statePopulations := map[string]int{
		"California":   39000000,
		"Texas":        200000,
		"New York":     300000000,
		"Pennsylvania": 120000000,
	}
	fmt.Println(statePopulations)
}

func statesMappingWithMakeFunc() {
	fmt.Printf("*** statesMappingWithMakeFunc ***\n")
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39000000,
		"Texas":        200000,
		"New York":     300000000,
		"Pennsylvania": 120000000,
	}
	fmt.Println(statePopulations)
	fmt.Println("Accessing key Texas", statePopulations["Texas"])
	fmt.Println("Accessing key Texas1", statePopulations["Texas1"])
	statePopulations["Georgia"] = 100000000
	fmt.Println(statePopulations)

	// delete key
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
}

func commaOkSyntax() {
	fmt.Printf("*** commaOkSyntax ***\n")
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39000000,
		"Texas":        200000,
		"New York":     300000000,
		"Pennsylvania": 120000000,
	}
	_, ok := statePopulations["Georgia"]
	fmt.Println(ok)

	_, ok2 := statePopulations["California"]
	fmt.Println(ok2)

	fmt.Println("Length of map: ", len(statePopulations))
}

func structDoctor() {
	fmt.Printf("*** structDoctor ***\n")
	type Doctor struct {
		number     int
		actorName  string
		companions []string
	}

	aDoctor := Doctor{
		number:    3,
		actorName: "Super Admin",
		companions: []string{
			"Dept Admin",
			"Regional Admin",
			"Local Admin",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions)
	fmt.Println(aDoctor.companions[2])
}

func anonymousStruct() {
	fmt.Printf("*** anonymousStruct ***\n")

	aDoctor := struct{ name string }{name: "Adam Kareem"}
	newDoctor := aDoctor
	newDoctor.name = "Sakura chan"
	latestDoc := &aDoctor
	fmt.Println("aDoctor ===> ", aDoctor)
	fmt.Println("newDoctor ====> ", newDoctor)
	fmt.Println("latestDoc ===> ", latestDoc)
	latestDoc.name = "Hinata chan"
	fmt.Println("aDoctor ===> ", aDoctor)
	fmt.Println("latestDoc ===> ", latestDoc)
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func structComposition() {
	fmt.Printf("*** structComposition ***\n")

	b := Bird{}
	b.Name = "Falcon"
	b.Origin = "Nigeria"
	b.SpeedKPH = 50
	b.CanFly = true
	fmt.Println("Bird ===> ", b)

	b2 := Bird{
		Animal:   Animal{Name: "Falconess", Origin: "Nigeria"},
		SpeedKPH: 40,
		CanFly:   true,
	}
	fmt.Println("Bird 2 ===> ", b2)
}

type Animal2 struct {
	Name   string `required max:100`
	Origin string
}

func structTags() {
	fmt.Printf("*** structTags ***\n")

	t := reflect.TypeOf(Animal2{})
	field, r := t.FieldByName("Name")
	fmt.Println("r ===> ", r)
	fmt.Println("Tag ===> ", field.Tag)
	fmt.Println("Name ===> ", field.Name)
	fmt.Println("Anonymous ===> ", field.Anonymous)
	fmt.Println("Index ===> ", field.Index)
}
