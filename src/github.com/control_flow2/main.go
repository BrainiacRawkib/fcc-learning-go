package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// deferOne
	fmt.Printf("\n")
	deferOne()

	// deferTwo
	fmt.Printf("\n")
	deferTwo()

	// googleRobot
	fmt.Printf("\n")
	//googleRobot()

	// panicOne
	fmt.Printf("\n")
	panicOne()

	// deferAndPanic
	fmt.Printf("\n")
	deferAndPanic()

	// runServerWithPanic
	//fmt.Printf("\n")
	//runServerWithPanic()
}

func deferOne() {
	fmt.Printf("*** deferOne ***\n")
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}

func deferTwo() {
	fmt.Printf("*** deferTwo ***\n")
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

func googleRobot() {
	fmt.Printf("*** googleRobot ***\n")
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func panicOne() {
	fmt.Printf("*** panicOne ***\n")
	//a, b := 1, 0
	//ans := a / b
	//panic("Zero division error!")
	//fmt.Println(ans)
}

func deferAndPanic() {
	fmt.Printf("*** deferAndPanic ***\n")
	fmt.Println("Starting...")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("Panicking...")
}

//func main() {
//	fmt.Printf("*** runServerWithPanic ***\n")
//	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
//		response.Write([]byte("An entire web development Course in less than 10 minutes!"))
//	})
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		panic(err.Error())
//	}
//}
