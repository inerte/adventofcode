package two

import "strings"

func EmptyKey() Key {
	return Key{"", false, false, false, false}
}

func DayTwoPartTwo(directions string) string {
	lines := strings.Split(directions, "\n")

	keypad := make(map[int]Key)
	keypad[0] = EmptyKey()
	keypad[1] = EmptyKey()
	keypad[2] = Key{"1", false, false, true, false}
	keypad[3] = EmptyKey()
	keypad[4] = EmptyKey()

	keypad[5] = EmptyKey()
	keypad[6] = Key{"2", false, true, true, false}
	keypad[7] = Key{"3", true, true, true, true}
	keypad[8] = Key{"4", false, false, true, true}
	keypad[9] = EmptyKey()

	keypad[10] = Key{"5", false, true, false, false}
	keypad[11] = Key{"6", true, true, true, true}
	keypad[12] = Key{"7", true, true, true, true}
	keypad[13] = Key{"8", true, true, true, true}
	keypad[14] = Key{"9", false, false, false, true}

	keypad[15] = EmptyKey()
	keypad[16] = Key{"A", true, true, false, false}
	keypad[17] = Key{"B", true, true, true, true}
	keypad[18] = Key{"C", true, false, false, true}
	keypad[19] = EmptyKey()

	keypad[20] = EmptyKey()
	keypad[21] = EmptyKey()
	keypad[22] = Key{"D", true, false, false, false}
	keypad[23] = EmptyKey()
	keypad[24] = EmptyKey()

	return KeypadCode(lines, keypad, 10)
}
