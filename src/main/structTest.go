package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Dict struct {
	data map[int]string
}

func newDict() *Dict {
	d := Dict{}
	d.data = map[int]string{}
	return &d
}

func main() {
	test2()
	test3()

	dict := newDict()
	dict.data[1] = "Hello"

	fmt.Println(dict)
}

func test2() {
	var person Person
	person = Person{"Bob", 20}
	person2 := Person{name: "Tom", age: 50}
	fmt.Println(person, person2)
}

func test3() {
	person := new(Person)
	fmt.Println(person)

	person.name = "Susan"
	fmt.Println(person)
}

func test1() {
	p := Person{}

	p.name = "typark"
	p.age = 32

	fmt.Println(p)
}
