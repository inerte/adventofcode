package two

import (
	"math"
	"strings"
)

type Key struct {
	code string
	U    bool
	R    bool
	D    bool
	L    bool
}

func KeypadCode(lines []string, keypad map[int]Key, currentKeypadIndex int) string {
	code := make([]string, len(lines))

	rowsQty := int(math.Sqrt(float64(len(keypad))))

	currentKey := keypad[currentKeypadIndex]

	for i, line := range lines {
		lineDirections := strings.Split(line, "")
		for _, direction := range lineDirections {
			if direction == "U" && currentKey.U == true {
				currentKeypadIndex = currentKeypadIndex - rowsQty
			}
			if direction == "R" && currentKey.R == true {
				currentKeypadIndex = currentKeypadIndex + 1
			}
			if direction == "D" && currentKey.D == true {
				currentKeypadIndex = currentKeypadIndex + rowsQty
			}
			if direction == "L" && currentKey.L == true {
				currentKeypadIndex = currentKeypadIndex - 1
			}
			currentKey = keypad[currentKeypadIndex]
		}
		code[i] = currentKey.code
	}
	return strings.Join(code, "")
}

func DayTwoPartOne(directions string) string {
	lines := strings.Split(directions, "\n")
	keypad := make(map[int]Key)

	keypad[0] = Key{"1", false, true, true, false}
	keypad[1] = Key{"2", false, true, true, true}
	keypad[2] = Key{"3", false, false, true, true}

	keypad[3] = Key{"4", true, true, true, false}
	keypad[4] = Key{"5", true, true, true, true}
	keypad[5] = Key{"6", true, false, true, true}

	keypad[6] = Key{"7", true, true, false, false}
	keypad[7] = Key{"8", true, true, false, true}
	keypad[8] = Key{"9", true, false, false, true}

	return KeypadCode(lines, keypad, 4)
}
