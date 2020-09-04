package main

import (
	"fmt"
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

func walkThroughSwitches(switches int) int {
	var lampOn int
	lamps := make([]bool, switches)

	for trip := 1; trip <= len(lamps); trip++ {
		for lampAt := trip - 1; lampAt < len(lamps); lampAt += trip {
			lamps[lampAt] = !lamps[lampAt]
		}
	}

	for _, lamp := range lamps {
		if lamp {
			lampOn++
		}
	}

	return lampOn
}

func main() {
	input1 := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Printf("===================================\n")
	fmt.Printf("Algorithm Number One (sockMerchant)\n")
	fmt.Printf("Sample Input\n%v\n%v\n", len(input1), input1)
	fmt.Printf("Sample Output\n%v\n", sockMerchant(input1))
	fmt.Printf("===================================\n")

	input2 := "UDDDUDUU"
	fmt.Printf("===================================\n")
	fmt.Printf("Algorithm Number Two (countingValleys)\n")
	fmt.Printf("Sample Input\n%v\n%v\n", len(input2), input2)
	fmt.Printf("Sample Output\n%v\n", countingValleys(input2))
	fmt.Printf("===================================\n")

	input3 := 1345679
	fmt.Printf("===================================\n")
	fmt.Printf("Algorithm Number Three (nominalLayers)\n")
	fmt.Printf("Sample Input\n%v\n", input3)
	fmt.Printf("Sample Output\n%v\n", nominalLayers(input3))
	fmt.Printf("===================================\n")

	input4 := 100
	fmt.Printf("===================================\n")
	fmt.Printf("Algorithm Number Four (walkThroughSwitches)\n")
	fmt.Printf("Sample Input\n%v\n", input4)
	fmt.Printf("Sample Output\n%v\n", walkThroughSwitches(input4))
	fmt.Printf("===================================\n")

}
