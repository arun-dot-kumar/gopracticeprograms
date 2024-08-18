package main

import "fmt"

/*
output:
    1
   1 1
  1 2 1
 1 3 3 1
1 4 6 4 1
*/

func factorial(n int) int {
	if n == 0 || n < 0 {
		return 1
	}
	return n * factorial(n-1)
}

func getBinoCoef(n int, k int) int {
	return factorial(n) / (factorial(n-k) * factorial(k))
}

func main() {
	for i := 0; i < 5; i++ {
		for j := 4; j > i; j-- {
			fmt.Printf(" ")
		}
		for p := 0; p < i+1; p++ {
			fmt.Printf("%d ", getBinoCoef(i, p))
		}
		fmt.Println()
	}
}
