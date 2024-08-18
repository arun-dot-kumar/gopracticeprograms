package main

import "fmt"

/*

output:
1
2 3
4 5 6
7 8 9 10

*/

func main() {
	var i = 1
	for j := 0; j < 4; j++ {
		for p := 0; p < j+1; p++ {
			fmt.Printf("%d ", i)
			i++
		}
		fmt.Println()
	}
}
