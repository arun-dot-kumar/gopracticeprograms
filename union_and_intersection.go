package main

import "fmt"

func main() {
	var arr1 = []int{1, 3, 4, 5, 7}
	var arr2 = []int{2, 3, 5, 6}
	var m = make(map[int]int)
	var m1 = make(map[int]int)
	var unionarr []int
	var intersectionarr []int
	for i := range arr1 {
		m[arr1[i]]++
		m1[arr1[i]]++
	}
	for i := range arr2 {
		m[arr2[i]]++
		if _, ok := m1[arr2[i]]; ok {
			intersectionarr = append(intersectionarr, arr2[i])
		}
	}
	for k, _ := range m {
		unionarr = append(unionarr, k)
	}
	fmt.Println("union of 2 array", unionarr)
	fmt.Println("intersection of 2 array", intersectionarr)

}
