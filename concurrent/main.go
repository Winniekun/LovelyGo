package main

import (
	"fmt"
	"time"
)

func channel() {
	ch := make(chan string)
	go func() {
		fmt.Println("协程工作")
		ch <- "工作完成"
	}()

	v := <- ch
	fmt.Println("从子协程中获取的数据", v)
}

func main() {
	go fmt.Println("我是维坤坤的 协程")
	fmt.Println("我是维坤坤的主线程")
	time.Sleep(time.Second)
	channel()
}
