package main

var x int

func main() {
	println("#######")
	println(&x, x)
	x = 100
	println(&x, x)
}
