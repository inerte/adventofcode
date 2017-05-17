package four

import (
	"math"
	"strings"
)

const abc = "abcdefghijklmnopqrstuvwxyz"
const abcLen = len(abc)

// From http://stackoverflow.com/a/36804185/14645
func rotateLetter(letter string, sectorID int) string {
	abcPos := strings.Index(abc, letter) + 1
	mod := math.Mod(float64(sectorID), float64(abcLen))
	rotateToIndex := abcPos + int(mod)
	if rotateToIndex > abcLen {
		rotateToIndex = rotateToIndex - abcLen
	}
	return string(abc[rotateToIndex-1])
}

func GetRealName(roomName string) string {
	encryptedName, sectorID, _ := GetRoomNameParts(roomName)
	realName := ""
	letters := strings.Split(strings.Join(encryptedName, " "), "")
	for _, letter := range letters {
		if letter == " " {
			realName = realName + " "
		} else {
			realName = realName + rotateLetter(letter, sectorID)
		}
	}
	return realName
}

func DayFourPartTwo(roomNames string) int {
	for _, roomName := range strings.Split(roomNames, "\n") {
		if GetRealName(roomName) == "northpole object storage" {
			_, sectorID, _ := GetRoomNameParts(roomName)
			return sectorID
		}
	}
	return 0
}
