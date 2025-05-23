package main

import "strconv"

func isPalindrome(x int) bool {
	original := strconv.Itoa(x)
	reversed := ""

	for _, c := range original {
		reversed = string(c) + reversed
	}

	return original == reversed
}