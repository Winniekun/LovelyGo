package main

func main() {
	// arr 的定义
	// 1 常规的定义
	arr := [3]int{1, 2, 3}
	println(arr[1])

	// 2 不指定容量大小
	nextArr := [...]int{4, 5, 6}
	println(nextArr[1])

	// 3 定一个空的数组
	var nullArr [3]int
	nullArr2 := [3]int{}
	nullArr2[2] = 10
	nullArr[1] = 1
	println(nullArr[0])
	println(nullArr2[2])
}
