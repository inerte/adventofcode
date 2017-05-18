package five

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
	"strings"
)

func passwordHasEmptyCharacters(password [8]string) bool {
	for _, character := range password {
		if character == "" {
			return true
		}
	}
	return false
}

func DayFivePartTwo(doorID string) string {
	var password [8]string
	i := 0
	// for passwordHasEmptyCharacters(password) {
	for passwordHasEmptyCharacters(password) != false {
		toHash := doorID + strconv.Itoa(i)
		h := md5.New()
		io.WriteString(h, toHash)
		encode := hex.EncodeToString(h.Sum(nil))
		if encode[:5] == "00000" {
			position, err := strconv.Atoi(string(encode[5]))
			if err == nil && position < 8 && password[position] == "" {
				character := string(encode[6])
				password[position] = character
			}
		}
		i++
	}
	return strings.Join(password[:], "")
}
