package main

import "fmt"

/*
output:
 * * * * * *
  * * * * *
   * * * *
    * * *
     * *
      *
*/

func main() {
	for i := 0; i <= 5; i++ {
		for j := i; j >= 0; j-- {
			fmt.Printf(" ")
		}
		for p := 5; p >= i; p-- {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
