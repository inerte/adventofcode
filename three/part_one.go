package three

import (
	"sort"
	"strconv"
	"strings"
)

func IsTriangle(angles string) bool {
	sidesAsString := strings.Split(angles, " ")
	sidesAsInts := make([]int, 0)

	for _, side := range sidesAsString {
		if side != "" {
			sideAsInt, _ := strconv.Atoi(side)
			sidesAsInts = append(sidesAsInts, sideAsInt)
		}
	}
	sort.Ints(sidesAsInts)

	return (sidesAsInts[0] + sidesAsInts[1]) > sidesAsInts[2]
}

func DayThreePartOne(input string) int {
	lines := strings.Split(input, "\n")
	trianglesQty := 0

	for _, line := range lines {
		cleanLine := strings.Trim(line, " ")
		if IsTriangle(cleanLine) {
			trianglesQty = trianglesQty + 1
		}
	}

	return trianglesQty
}
