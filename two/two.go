package two

import (
	"math"
	"strconv"
	"strings"
)

func DayTwo(directions string) string {
	lines := strings.Split(directions, "\n")
	codeSize := len(lines)
	code := make([]string, codeSize)

	currentIndex := 5
	for i, line := range lines {
		lineDirections := strings.Split(line, "")
		for _, direction := range lineDirections {
			if direction == "U" && currentIndex > 3 {
				currentIndex = currentIndex - 3
			}
			if direction == "R" && math.Mod(float64(currentIndex), 3) != 0 {
				currentIndex = currentIndex + 1
			}
			if direction == "D" && currentIndex < 7 {
				currentIndex = currentIndex + 3
			}
			if direction == "L" && math.Mod(float64(currentIndex), 3) != 1 {
				currentIndex = currentIndex - 1
			}
		}
		code[i] = strconv.Itoa(currentIndex)
	}
	return strings.Join(code, "")
}
