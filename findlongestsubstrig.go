package main

import "fmt"

func kthsmallest(arr string, k int) {
	var a = make(map[string]int)
	var start, end int
	var i int
	var substring = make(map[int][]string)
	for i = range arr {
		a[string(arr[i])] = i
		if len(a) > k {
			for len(a) > k {
				if end < i-start {
					substring = make(map[int][]string)
					substring[i-start] = append(substring[i-start], arr[start:i])
					end = i - start
				}
				if a[string(arr[start])] != start {
					start++
				} else {
					delete(a, string(arr[start]))
					start++
				}
			}
		}
	}
	if end < i-start {
		substring = make(map[int][]string)
		substring[i-start] = append(substring[i-start], arr[start:i+1])
		end = i - start
	}
	fmt.Println(substring[end])
}

func main() {
	kthsmallest("abcbdbdbbdcdabd", 2)
}
