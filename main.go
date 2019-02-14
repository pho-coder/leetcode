package main

func main() {
	a := "abcd"
	b := []byte(a)
	println(string(b[:4]))
}
