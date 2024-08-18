package main

import "fmt"

/*
output:
    *
   * *
  * * *
 * * * *
 * * * *
  * * *
   * *
    *
*/

func main() {
	for i := 0; i < 4; i++ {
		for p := 0; p < 4-i; p++ {
			fmt.Printf(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
	for i := 0; i < 4; i++ {
		for p := 0; p <= i; p++ {
			fmt.Printf(" ")
		}
		for j := 3; j >= i; j-- {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
