package main

import (
	"fmt"
	"primespace/testlib"
)

func main() {
	fmt.Println("Package..")

	name := testlib.GetMusic("Adele")

	fmt.Println(name)

}
