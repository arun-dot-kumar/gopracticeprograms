package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 7}
	target := 6
	var sum int
	var res [][]int
	var startingindex int
	for i := range a {
		sum += a[i]
		if sum == target {
			var temp []int
			temp = append(temp, a[startingindex:(i+1)]...)
			res = append(res, temp)
			for sum < target {
				sum -= a[startingindex]
				startingindex++
			}
		}
	}

	sum = 0
	var m1 = make(map[int]int)
	for i := range a {
		if _, ok := m1[target-a[i]]; ok {
			var temp []int
			temp = append(temp, a[m1[target-a[i]]])
			temp = append(temp, a[i])
			res = append(res, temp)
		}
		m1[a[i]] = i
	}
	fmt.Println(res)
}
