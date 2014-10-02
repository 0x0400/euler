package main

import (
	"fmt"
	"strconv"
)

func count(divisor int) int {
	devident, remainder := 10, 1
	remainders := map[int]int{1: 0}
	quotient := ""
	for {
		if devident < divisor {
			devident *= 10
			quotient += "0"
		} else {
			remainder = devident % divisor
			quotient += strconv.Itoa(devident / divisor)
			if remainder == 0 {
				return 0
			}
			if idx, exists := remainders[remainder]; exists {
				return len(quotient) - idx
			} else {
				remainders[remainder] = len(quotient) - 1
				devident = remainder * 10
			}
		}
	}

}

func main() {
	max_cycle_len, max_denominator := 0, 0
	for denominator := 2; denominator < 1000; denominator++ {
		cycle_len := count(denominator)
		fmt.Println(denominator, cycle_len)
		if cycle_len > max_cycle_len {
			max_cycle_len, max_denominator = cycle_len, denominator
		}
	}
	fmt.Println(max_denominator, max_cycle_len)
}
