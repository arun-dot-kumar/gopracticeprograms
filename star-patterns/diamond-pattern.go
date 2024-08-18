package main

import "fmt"

/*
output:

*/
func main() {
	for i := 0; i < 4; i++ {
		for j := 3; j > i; j-- {
			fmt.Printf("-")
		}
		for p := 0; p <= i; p++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
	for i := 0; i < 4; i++ {
		for j := 3; j > i; j-- {
			fmt.Printf("-")
		}
		for p := 0; p <= i; p++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
