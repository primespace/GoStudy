package main

func main() {

	testAnonymous()

	testFunctionParam()

	testFunctionUsingType()
}

type calculator func(int, int) int

func calc2(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

func testFunctionUsingType() {

}

func testFunctionParam() {
	add := func(i int, j int) int {
		return (i + j)
	}

	r1 := calc(add, 10, 20)
	println(r1)
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func testAnonymous() {
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5)
	println(result)
}
