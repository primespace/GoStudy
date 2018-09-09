package main

const c int = 10
const s string = "Hi"

const (
	Visa   = "Visa"
	Master = "MasterCard"
	Amex   = "American Express"
)

const (
	Apple = iota
	Grape
	Orange
)

func main() {
	testVariables()

	// Constants
	println(c)
	println(s)
	println(Visa, Master, Amex)
	println(Apple, Grape, Orange)

}

func testVariables() {
	var a int

	a = 100

	var i, j, k int = 1, 2, 3

	var m int = 1234
	// var a float32 = 10. // error

	println(a)
	println(i, j, k)
	println(m)
}
