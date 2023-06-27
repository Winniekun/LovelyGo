package main

func plusOne(digits []int) []int {
	// 思路
	// 切片 存储结果
	n := len(digits)
	arr := make([]int, n+1)
	carry := 0
	base := 1
	for i := n - 1; i >= 0; i-- {
		cur := digits[i] + carry + base
		arr[i+1] = cur % 10
		carry = cur / 10
		base = 0
	}
	if carry > 0 {
		arr[0] = carry
	}
	return arr
}
