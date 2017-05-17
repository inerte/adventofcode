package three

import (
	"math"
	"strconv"
	"strings"
)

func columnNowHasTriangle(column []int, sideAsInt int, trianglesQty int) ([]int, int) {
	column = append(column, sideAsInt)
	if len(column) == 3 && IsTriangle(column) {
		trianglesQty = trianglesQty + 1
	}
	return column, trianglesQty
}

func DayThreePartTwo(input string) int {
	lines := strings.Split(input, "\n")
	trianglesQty := 0

	columnA := make([]int, 0)
	columnB := make([]int, 0)
	columnC := make([]int, 0)

	for i, line := range lines {
		// Reset columns every 3 lines
		iMod := int(math.Mod(float64(i), 3))
		if iMod == 0 {
			columnA = columnA[:0]
			columnB = columnB[:0]
			columnC = columnC[:0]
		}

		sidesAsString := strings.Split(strings.Trim(line, " "), " ")
		j := 0
		for _, sideAsString := range sidesAsString {
			if sideAsString != "" {
				jMod := int(math.Mod(float64(j), 3))
				sideAsInt, _ := strconv.Atoi(sideAsString)

				if jMod == 0 {
					columnA, trianglesQty = columnNowHasTriangle(columnA, sideAsInt, trianglesQty)
				}
				if jMod == 1 {
					columnB, trianglesQty = columnNowHasTriangle(columnB, sideAsInt, trianglesQty)
				}
				if jMod == 2 {
					columnC, trianglesQty = columnNowHasTriangle(columnC, sideAsInt, trianglesQty)
				}
				j = j + 1
			}
		}
	}

	return trianglesQty
}
