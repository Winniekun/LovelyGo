package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var i type = expression
	var i int = 10
	fmt.Println(i)
	// 因为具有自动的类型推导，可以让go自己进行推到
	var j = 10
	fmt.Println(j)

	// 多个变量的定义 可进行合并
	var (
		a = 10
		b = "hello"
	)
	println(a)
	println(b)

	// 测试各种类型的0值
	var in int
	var fl float32
	var str string
	var bo bool
	fmt.Println("int default: ", in, " float default : ", fl, " string default : ", str,
		" boolean default: ", bo)
	pi :=&i
	println(pi)
	i = 20
	println(pi)

	// 类型转换
	// string -> int
	i2s := strconv.Itoa(i)
	println(i2s)
	// int -> string
	s2i, _ := strconv.Atoi(i2s)
	println(s2i)

}
