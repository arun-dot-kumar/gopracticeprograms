package main

import "fmt"

func checkanagram(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	var m = make(map[string]int)
	for i := range str1 {
		m[string(str1[i])]++
	}
	for i := range str2 {
		m[string(str2[i])]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func substringofanotherstring(str1 string, str2 string) {
	windowsize := len(str2)
	var res []string
	for i := 0; i < len(str1)-windowsize+1; i++ {
		if checkanagram(str2, str1[i:i+windowsize]) {
			res = append(res, str1[i:i+windowsize])
		}
	}
	fmt.Println(res)
}

func main() {
	substringofanotherstring("XYYZXZYZXXYZ", "XYZ")
}
