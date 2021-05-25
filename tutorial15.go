package main

import "fmt"

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}

	var s []int = x[1:3]

	fmt.Println(len(s))

	fmt.Println(s[:cap(s)])

	fmt.Println("*********")

	var a []int = []int{5, 6, 7, 8, 9}
	//fmt.Println(cap(a[:3]))

	b := append(a, 10)

	fmt.Println(b)

	fmt.Println("*********")

	m := make([]int, 5)

	fmt.Printf("%T", m)

}
