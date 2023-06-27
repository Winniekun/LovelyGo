package main

import (
	"fmt"
	"sync"
	"time"
)

func multiTask() {
	firstCh := make(chan string)
	secondCh := make(chan string)
	thirdCh := make(chan string)

	go func() {
		firstCh <- downloadFile("firstCh")
	}()

	go func() {
		secondCh <- downloadFile("secondCh")
	}()

	go func() {
		thirdCh <- downloadFile("thirdCh")
	}()

	select {
	case firstPath := <-firstCh:
		fmt.Println(firstPath)
	case secondPath := <-secondCh:
		fmt.Println(secondPath)
	case thirdPath := <-thirdCh:
		fmt.Println(thirdPath)
	}
}

func downloadFile(chanName string) string {

	//模拟下载文件,可以自己随机time.Sleep点时间试试

	time.Sleep(time.Second)

	return chanName + ":filePath"

}

func channelTest() {
	cacheCh := make(chan int, 5)
	cacheCh <- 2
	cacheCh <- 3
	fmt.Println("容量为: ", cap(cacheCh), ", 元素个数: ", len(cacheCh))
}

var sum = 0
func multiAdd()  {

	for i := 0; i < 10000; i++ {
		go add(1)
	}

}

var mutex sync.Mutex
func add(num int)  {
	//mutex.Lock()
	//defer mutex.Unlock()
	sum += num
}

func main() {
	//multiTask()
	//channelTest()
	multiAdd()
	time.Sleep( time.Second)
	println(sum)
}
