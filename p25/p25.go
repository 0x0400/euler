package main

import (
	"fmt"
	"math/big"
)

func main() {
	fab1 := big.NewInt(1)
	fab := big.NewInt(1)
	idx := 2
	for len(fab.String()) < 1000 {
		fab1, fab = big.NewInt(0).Set(fab), fab.Add(fab1, fab)
		idx += 1
	}
	fmt.Println(fab)
	fmt.Println(idx)
}
