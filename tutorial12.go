package main

import "fmt"

func main() {

	x := 0

	for x < 5 {

		fmt.Println(x)
		x++

	}

	fmt.Println("*********")

	for x := 0; x <= 5; x++ {

		fmt.Println(x)
	}

	fmt.Println("*********")

	for x := 0; x <= 10000; x++ {

		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			break
		}

	}

}
