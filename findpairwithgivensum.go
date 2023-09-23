package main

import "fmt"

func findSumPair(arr []int, sum int) {
	mapping := make(map[int]int)
	var a [][]int
	for i := 0; i < len(arr); i++ {
		if _, ok := mapping[sum-arr[i]]; !ok {
			mapping[arr[i]] = i
		} else {
			a = append(a, []int{sum - arr[i], arr[i]})
		}
	}
	fmt.Println(a)
}

func main() {
	findSumPair([]int{8, 7, 2, 5, 3, 1}, 10)
	findSumPair([]int{4, 3, 6, 7, 8, 1, 9}, 100)
}
