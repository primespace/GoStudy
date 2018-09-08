package main

import "fmt"

func main() {

	test1()

	test2()

	testAppend()

	testCopy()

}

func testCopy() {
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)

	copy(target, source)
	fmt.Println(target)
	println(len(target), cap(target))
}

func testAppend() {
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB...)

	fmt.Println(sliceA)
}

func test1() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	s := make([]int, 5, 10)
	println(len(s), cap(s))
}

func test2() {
	s := []int{0, 1, 2, 3, 4, 5}
	s = s[2:5] // 2, 3, 4
	s = s[1:]  // 3, 4
	fmt.Println(s)

	sliceA := make([]int, 0, 3)

	for i := 1; i < 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA)
}
