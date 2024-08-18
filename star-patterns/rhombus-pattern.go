package main

import "fmt"

/*
output:

 * * * *
  * * * *
   * * * *
    * * * *
     * * * *
      * * * *
*/

func main() {
	for i := 0; i <= 5; i++ {
		for j := i; j >= 0; j-- {
			fmt.Printf(" ")
		}
		for p := 0; p < 4; p++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
