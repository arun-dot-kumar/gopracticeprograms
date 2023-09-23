package main

import "fmt"

func substring_with_distinct(str string) {
	var charcterfreq = make(map[string]int)
	var charcterpos = make(map[string]int)
	var res string
	var start int
	var i int
	for i = range str {
		if _, ok := charcterfreq[string(str[i])]; !ok {
			charcterfreq[string(str[i])]++
			charcterpos[string(str[i])] = i
		} else {
			charcterfreq = make(map[string]int)
			if len(res) < i-start {
				res = str[start:i]
			}
			start = charcterpos[string(str[i])] + 1
			charcterpos[string(str[i])] = i
			charcterfreq[string(str[i])]++
		}
	}
	if len(res) < i-start {
		res = str[start : i+1]
	}
	fmt.Println(res)
}

func main() {
	/*
	Input:  findlongestsubstring

	Output: The longest substring with all distinct characters is dlongest or ubstring


	Input:  longestsubstr

	Output: The longest substring with all distinct characters is longest


	Input:  abbcdafeegh

	Output: The longest substring with all distinct characters is bcdafe


	Input:  aaaaaa

	Output: The longest substring with all distinct characters is a
	*/
	substring_with_distinct("abbcdafeegh")
}
