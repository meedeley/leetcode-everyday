package main

import "fmt"

func main() {
	value := "MENINGGAL"
	reversed := ""

	for _, s := range value {
		reversed = string(s) + reversed
	}

	fmt.Println(value == reversed)
}

// func isPalindrome(x int) bool {
 
// 	for i := 0; i < len(x); i++ {

// 	}
// }