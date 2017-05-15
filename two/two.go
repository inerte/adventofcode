package two

import (
	"math"
	"strings"
)

func DayTwo(directions string) [4]float64 {
	var code [4]float64

	lines := strings.Split(directions, "\n")
	currentIndex := float64(5)
	for i, line := range lines {
		lineDirections := strings.Split(line, "")
		for _, direction := range lineDirections {
			if direction == "U" && currentIndex > 3 {
				currentIndex = currentIndex - 3
			}
			if direction == "R" && math.Mod(currentIndex, 3) != 0 {
				currentIndex = currentIndex + 1
			}
			if direction == "D" && currentIndex < 7 {
				currentIndex = currentIndex + 3
			}
			if direction == "L" && math.Mod(currentIndex, 3) != 1 {
				currentIndex = currentIndex - 1
			}
		}
		code[i] = currentIndex
	}
	return code
}
