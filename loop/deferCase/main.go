package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("延迟执行: ", time.Now())
	time.Sleep(time.Second)
	fmt.Println("开始执行: ", time.Now())

}
