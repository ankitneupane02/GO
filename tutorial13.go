package main

import "fmt"

func main() {
	ans := 10

	switch ans {
	case 1:
		fmt.Println("1.one")
		fmt.Println("2.one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("you are ")
	default:
		fmt.Println("not a case")

	}

}
