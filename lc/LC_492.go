package main

import "math"

func constructRectangle(area int) []int {
	arr := make([]int, 2)
	w := int(math.Sqrt(float64(area)))
	for area % w > 0 {
		w--
	}
	arr[0] = w
	arr[1] = area / w
	return arr
}
