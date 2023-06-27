package main

import (
	"fmt"
	"runtime"
	"time"
)

func switchWithNoCondition() {
	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	
}

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s. \n", os)
	}
	switchWithNoCondition()
}