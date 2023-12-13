package main

import (
	"fmt"
)

func pairsOfSocks(arr []int) int {

	colorsCount := map[int]int{}

	// 10, 20, 20, 10, 10, 30, 50, 10, 20
	for _, color := range arr {
		colorsCount[color]++
	}

	// map[10:4, 20:3, 30:1, 50:2]
	pairs := 0
	for _, count := range colorsCount {
		pairs = pairs + count/2
	}

	return pairs
}

func main() {
	arrSocks := []int{10, 20, 20, 10, 10, 30, 50, 10, 20, 50}
	resultOfPairs := pairsOfSocks(arrSocks)
	fmt.Println(resultOfPairs)
}
