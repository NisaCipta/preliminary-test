package main

import "fmt"

func turnOnTheLight(number int) int {
	lightOn := 0

	for lamp := 1; lamp <= number; lamp++ {
		factor := 0
		for turnOn := 1; turnOn <= lamp; turnOn++ {
			if lamp%turnOn == 0 {
				factor++
			}
		}
		if factor%2 == 1 {
			lightOn++
		}
	}
	return lightOn
}

func main() {
	fmt.Println(turnOnTheLight(100))
}
