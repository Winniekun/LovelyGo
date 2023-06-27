package main

import (
	"fmt"
	"os"
)

func main() {
	s, step := "", ""
	for _, arg := range os.Args[1:] {
		s += step + arg
		step = " "
	}
	//fmt.Println(s)
	fmt.Println(os.Args[1:])
}

func stringJoin() {
}
