package main

import (
	"fmt"
	"math"
)

const (
	MIN_ABUNDANT = 12
	MIN_NUMBER   = 24
	MAX_NUMBER   = 28123
)

func properDivisors(number int) []int {
	if number == 1 {
		return []int{}
	}
	vals := []int{1}
	sqrt := int(math.Sqrt(float64(number)))
	for idx := 2; idx <= sqrt; idx++ {
		if number%idx == 0 {
			vals = append(vals, idx)
			another := number / idx
			if another != idx {
				vals = append(vals, another)
			}
		}
	}
	return vals
}

func sum(slice []int) int {
	rst := 0
	for _, val := range slice {
		rst += val
	}
	return rst
}

func isAbundant(number int) bool {
	divisors := properDivisors(number)
	return sum(divisors) > number
}

func main() {
	abundans := []int{}
	var numbers [MAX_NUMBER]int
	for idx := len(numbers); idx > 0; idx-- {
		numbers[idx-1] = idx
	}
	for idx := 1; idx <= MAX_NUMBER; idx++ {
		if isAbundant(idx) {
			abundans = append(abundans, idx)
			for _, abundant := range abundans {
				if abundant+idx-1 < MAX_NUMBER {
					numbers[abundant+idx-1] = 0
				}
			}
		}
	}
	fmt.Println(sum(numbers[:]))
}
