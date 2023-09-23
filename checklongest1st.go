package main

import "fmt"

func substring_with_distinct(arr []int) {
	var a = make(map[int]int)
	var start int
	var i int
	for i = range arr {
		if arr[i] == 0 {
			if _, ok := a[arr[i]]; ok {
				start = i
			} else {
				a[arr[i]]++
			}
		}
	}
	fmt.Println(start)
}

func main() {
	/*
      For example, consider array { 0, 0, 1, 0, 1, 1, 1, 0, 1, 1 }. The index to be replaced is 7 to get a continuous sequence of length 6 containing all 1’s.
	 */
	substring_with_distinct([]int{0, 0, 0, 1, 0, 0, 0, 1, 0, 1})
}
