package main

import (
	"fmt"
	"strings"
)

/*

output:
* * * * *
*       *
*       *
*       *
*       *
*       *
* * * * *

*/

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("* ")
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Printf("* ")
		//for j := 0; j < 3; j++ {
		//	fmt.Printf("  ")
		//}
		fmt.Printf(strings.Repeat("  ", 3))
		fmt.Printf("* ")
		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("* ")
	}
}
