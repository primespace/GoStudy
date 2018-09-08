package main

func main() {

	test1()

	test2()

}

func test2() {

	println("#### test2")

	var a = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for _, row := range a {
		for _, n := range row {
			println(n)
		}
	}
}

func test1() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1])

	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3}

	for index, n := range a1 {
		println(index, n)
	}

	for index, n := range a3 {
		println(index, n)
	}
}
