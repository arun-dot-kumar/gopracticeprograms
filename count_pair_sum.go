package main

import "fmt"

func main() {
	var arr1 = []int{1, 5, 7, -1}
	k := 6
	var m = make(map[int]int)
	var count int
	for i := range arr1 {
		m[arr1[i]]++
	}
	var m1 = make(map[int]int)
	var res [][]int
	for i := range arr1 {
		if _, ok := m[k-arr1[i]]; ok {
			m1[arr1[i]] = 1
			if _, ok1 := m1[k-arr1[i]]; !ok1 {
				count += m[k-arr1[i]]
				for j := 0; j < m[k-arr1[i]]; j++ {
					res = append(res, []int{arr1[i], k - arr1[i]})
				}
			}
		}
	}
	fmt.Println(count)
	fmt.Println(res)

}
