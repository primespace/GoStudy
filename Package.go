package main

import (
	"fmt"

	"./testlib"
)

func main() {
	fmt.Println("Package..")

	name := testlib.GetMusic("Adele")

	fmt.Println(name)
}
