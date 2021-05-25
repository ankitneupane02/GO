package main

import "fmt"

func main() {

	var x int = 50

	var y int = 20

	age := 19

	if y > x {
		fmt.Println("y is grater than x")

	} else {
		fmt.Println("x is greater")
	}

	if age >= 18 {

		fmt.Println("You can ride alone")

	} else if age >= 14 {
		fmt.Println("You can ride with a parent")

	} else {
		fmt.Println("You cannot ride!")
	}

}
