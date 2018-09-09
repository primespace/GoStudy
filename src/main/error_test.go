package main

import (
	"fmt"
	"log"
)

type MyError struct {
	s string
}

func (e MyError) Error() string {
	return e.s
}

func someFunc() error {
	return MyError{"someFunc error."}
}

func main() {

	err := someFunc()
	fmt.Println(err)

	switch err.(type) {
	default:
		println("ok")
	case MyError:
		log.Print("MyError.. ")
	case error:
		log.Fatal(err.Error())
	}
}
