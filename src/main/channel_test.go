package main

import "fmt"

func main() {

	test2()

}

func test2() {
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i += 1 {

			fmt.Println(i)
		}
		done <- true
	}()

	<-done
}

func test1() {
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	var i int
	i = <-ch
	println(i)
}
