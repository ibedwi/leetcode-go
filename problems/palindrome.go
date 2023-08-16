package problems

import (
	"fmt"
	"strconv"
)

func IsPalindrome(x int) bool {

	isPalindrome := true
	stringOfX := strconv.Itoa(x)

	fmt.Print("asdfasdf"[0])
	if string(x) == "1" {

	}
	runeOfX := []rune(stringOfX)
	lengthOfX := len(stringOfX)

	for i := 0; i < (lengthOfX / 2); i++ {
		if runeOfX[(lengthOfX-1)-i] != runeOfX[i] {
			isPalindrome = false
			break
		}
	}

	return isPalindrome
}
