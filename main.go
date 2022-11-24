package main

import (
	"fmt"
	"math"
)

func main() {
	list := []int{2, 5, 1, 2, 3, 4, 7, 7, 6}
	var left, leftMax, rightMax int
	var ans float64
	right := len(list) - 1
	for ok := true; ok; ok = left < right {
		if list[left] > leftMax {
			leftMax = list[left]
		}
		if list[right] > rightMax {
			rightMax = list[right]
		}
		if leftMax < rightMax {
			ans += math.Max(0, float64(leftMax-list[left]))
			left++
		} else {
			ans += math.Max(0, float64(rightMax-list[right]))
			right--
		}
	}

	fmt.Println(fmt.Sprintf("Output: %v", ans))
}
