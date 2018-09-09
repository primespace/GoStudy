package main

// Pass by reference

func say(msg *string) {
	println(*msg)
	*msg = "Hello World !"
}

// Variadic function
func say2(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func sum(nums ...int) (count int, total int) {

	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}

func main() {

	msg := "Hello"
	say(&msg)
	println(msg)

	say2("Hello", "World", "!")
	count, total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	println(count, total)
}
