package main

import "fmt"

func main() {
	var a = []int{7, 9, 8, 1, 6}
	var k = 4
	var maxsum int
	for i := 0; i < 4; i++ {
		maxsum += a[i]
	}
	tempsum := maxsum
	for i := 0; i <= len(a)-1-k; i++ {
		tempsum = tempsum - a[i] + a[i+k]
		if tempsum > maxsum {
			maxsum = tempsum
		}
	}
	fmt.Println(maxsum)
}
