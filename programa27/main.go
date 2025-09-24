package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1{
		if runes[i] != runes[j]{
			return false
		}
	}
	return true
}

func main() {
	x := []int{1,2,2}
	str:= ""
	for _, num := range x {
		str += strconv.Itoa(num)
	}
	fmt.Println(isPalindrome(str))
}