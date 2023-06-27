package main

func majorityElement(nums []int) []int {
	hashMap := make(map[int]int)
	var arr []int
	n := len(nums)
	for _, num := range nums {
		hashMap[num] += 1
	}

	for k, v := range hashMap {
		if v > n%3 {
			arr = append(arr, k)
		}
	}
	return arr
}
