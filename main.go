package main

import "strings"

func sockMerchant(colors []int) int {
	var pairSocks int

	colorMap := make(map[int]int)
	for i := 0; i < len(colors); i++ {
		if _, found := colorMap[colors[i]]; found {
			colorMap[colors[i]]++
			if colorMap[colors[i]]%2 == 0 {
				pairSocks++
			}
		} else {
			colorMap[colors[i]] = 1
		}
	}

	return pairSocks
}

func countingValleys(steps string) int {
	var counted int

	arrStr := strings.Split(steps, "")
	stepMap := map[string]int{"U": 0, "D": 0}
	seaLevel := true

	for i := 0; i < len(arrStr); i++ {
		stepMap[arrStr[i]]++
		currState := stepMap["U"] - stepMap["D"]

		if currState < 0 && seaLevel {
			counted++
		}

		if currState != 0 {
			seaLevel = false
		} else {
			seaLevel = true
		}
	}

	return counted
}

func main() {

}
