package main

import "fmt"

func countingValleys(str string) int {
	up := 0
	down := 0
	valley := false

	for _, step := range str {
		if string(step) == "U" {
			up++
		}

		if string(step) == "D" {
			up--
		}

		if up < 0 && !valley {
			valley = true
		}

		if up == 0 && valley {
			down++
			valley = false
		}
	}
	return down
}

func main() {
	arrHiker := "UDDDUDUU"
	resultOfHiker := countingValleys(arrHiker)
	fmt.Println(resultOfHiker)
}
