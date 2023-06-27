package main

// defer 按照栈的顺序执行，先出现的后执行
func main() {
	defer println("坤坤")
	defer println("是")
	defer println("我")
}
