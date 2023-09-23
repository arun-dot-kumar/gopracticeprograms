package main

import "fmt"

func count(arr []int, mid int) int {
	var c int
	for i := range arr {
		if arr[i] <= mid {
			c++
		}
	}
	return c

}

func kthsmallest(arr []int, k int) int {
	var max, min int
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	for min < max {
		mid := (max + min) / 2

		if count(arr, mid) < k {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return min
}

func main() {
	fmt.Println(kthsmallest([]int{8, 7, 2, 5, 3, 1}, 3))
}
