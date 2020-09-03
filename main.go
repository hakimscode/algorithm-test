package main

func algorithmNumberOne(colors []int) int {
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

func main() {

}
