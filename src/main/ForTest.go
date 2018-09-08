package main

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	n := 1
	for n < 100 {
		n *= 2

	}

	println(n)

	// for {
	// 	println("Infinite loop")
	// }

	names := []string{"Hello", "World", "!"}

	for index, name := range names {
		println(index, name)
	}

}
