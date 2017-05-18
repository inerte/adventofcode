package five

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
	"strings"
)

func GetPasswordForDoorID(doorID string) string {
	password := make([]string, 0)
	i := 0
	for len(password) < 8 {
		toHash := doorID + strconv.Itoa(i)
		h := md5.New()
		io.WriteString(h, toHash)
		encode := hex.EncodeToString(h.Sum(nil))
		if encode[:5] == "00000" {
			password = append(password, string(encode[5]))
		}
		i++
	}
	return strings.Join(password, "")
}
