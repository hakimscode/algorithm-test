package main

import (
	"math"
	"strconv"
	"strings"
)

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

func nominalLayers(nominal int) []int {
	var layers []int

	nominalArr := strings.Split(strconv.Itoa(nominal), "")

	for ind, eachNominal := range nominalArr {
		multiplier := math.Pow(10, float64(len(nominalArr)-1-ind))
		nominalVal, _ := strconv.Atoi(eachNominal)
		layers = append(layers, nominalVal*int(multiplier))
	}

	return layers
}

func main() {

}
