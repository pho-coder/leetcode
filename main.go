package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{11, 12}
	c := append(a, b...)
	fmt.Println(c)
}
