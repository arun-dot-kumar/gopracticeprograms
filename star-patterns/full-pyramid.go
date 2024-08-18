package main

import "fmt"

/*
output:
     *
    * *
   * * *
  * * * *
 * * * * *x
* * * * * *
*/

func main() {
	for i := 0; i <= 5; i++ {
		for j := 5; j > i; j-- {
			fmt.Printf(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
