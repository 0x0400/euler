package main

import (
	"fmt"
	"strconv"
)

func Remove(slice []int, start int, end int) []int {
	if start < 0 || end > len(slice) {
		panic("Out of range")
	}
	return append(slice[:start], slice[end:]...)
}

func main() {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	left := 1000000
	permutation_size := 1
	for idx := 1; idx <= len(digits); idx++ {
		permutation_size *= idx
	}
	permutation := ""
	for left > 0 && len(digits) > 0 {
		permutation_size /= len(digits)
		times := left / permutation_size
		mod := left % permutation_size
		if times > len(digits) {
			panic("no choice")
		}
		if mod == 0 {
			permutation += strconv.Itoa(digits[times-1])
			digits = Remove(digits, times-1, times)
			for idx := len(digits) - 1; idx >= 0; idx-- {
				permutation += strconv.Itoa(digits[idx])
			}
			digits = []int{}
		} else {
			permutation += strconv.Itoa(digits[times])
			digits = Remove(digits, times, times+1)
		}
		left = mod
	}
	fmt.Println(permutation)
}
