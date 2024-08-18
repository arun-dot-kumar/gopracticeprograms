package main

import "fmt"

/*
* * * * * *
  * * * * *
    * * * *
      * * *
        * *
          *
*/

func main() {
	for i := 5; i >= 0; i-- {
		for j := 0; j < 5-i; j++ {
			fmt.Printf("  ")
		}
		for p := 0; p <= i; p++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
