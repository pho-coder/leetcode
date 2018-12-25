package main

import (
	"fmt"
)

func main() {
	amap := map[int]string{1: "a", 2: "b"}
	val, ok := amap[3]
	fmt.Println(val)
	fmt.Println(ok)
	fmt.Println(len(amap))
}
