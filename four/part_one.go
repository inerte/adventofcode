package four

import (
	"sort"
	"strings"
)

var letterCount map[string]int

type Pair struct {
	Letter string
	Count  int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Count == p[j].Count {
		return p[i].Letter > p[j].Letter
	} else {
		return p[i].Count < p[j].Count
	}
}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Based on http://stackoverflow.com/a/18695740/14645
func FiveMostCommonLettersInLetterCount(letterCount map[string]int) string {
	pl := make(PairList, len(letterCount))
	i := 0
	for k, v := range letterCount {
		pl[i] = Pair{k, v}
		i++
	}
	// Custom sort by count, and if count is tied, by alphabetical order
	sort.Sort(sort.Reverse(pl))
	mostCommon := make([]string, 5)
	for i, p := range pl[0:5] {
		mostCommon[i] = p.Letter
	}
	return strings.Join(mostCommon, "")
}

func isRealRoom(roomName string) bool {
	letterCount = make(map[string]int)

	nameParts := strings.Split(roomName, "-")

	// From https://github.com/golang/go/wiki/SliceTricks#pop
	sectorIdAndChecksum, encryptedName := nameParts[len(nameParts)-1], nameParts[:len(nameParts)-1]

	for _, letterGroup := range encryptedName {
		letters := strings.Split(letterGroup, "")
		for _, letter := range letters {
			_, ok := letterCount[letter]
			if ok {
				letterCount[letter] = letterCount[letter] + 1
			} else {
				letterCount[letter] = 1
			}
		}
	}
	common := FiveMostCommonLettersInLetterCount(letterCount)
	// sectorId, _ := strconv.Atoi(sectorIdAndChecksum[:3])
	checksum := sectorIdAndChecksum[4 : len(sectorIdAndChecksum)-1]

	return common == checksum
}
