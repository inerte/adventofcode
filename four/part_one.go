package four

import (
	"sort"
	"strconv"
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
	}
	return p[i].Count < p[j].Count
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
	nameParts := strings.Split(roomName, "-")

	// From https://github.com/golang/go/wiki/SliceTricks#pop
	sectorIDAndChecksum, encryptedName := nameParts[len(nameParts)-1], nameParts[:len(nameParts)-1]
	checksum := sectorIDAndChecksum[4 : len(sectorIDAndChecksum)-1]

	return encryptedNameMatchesChecksum(encryptedName, checksum)
}

func encryptedNameMatchesChecksum(encryptedName []string, checksum string) bool {
	letterCount = make(map[string]int)

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

	return common == checksum
}

func getSectorIDIfIsRealRoom(roomName string) int {
	nameParts := strings.Split(roomName, "-")

	// From https://github.com/golang/go/wiki/SliceTricks#pop
	sectorIDAndChecksum, encryptedName := nameParts[len(nameParts)-1], nameParts[:len(nameParts)-1]
	sectorID, _ := strconv.Atoi(sectorIDAndChecksum[:3])
	checksum := sectorIDAndChecksum[4 : len(sectorIDAndChecksum)-1]

	if encryptedNameMatchesChecksum(encryptedName, checksum) {
		return sectorID
	}
	return 0
}

func DayFourPartOne(roomNames string) int {
	sectorIDSum := 0
	for _, roomName := range strings.Split(roomNames, "\n") {
		sectorIDSum = sectorIDSum + getSectorIDIfIsRealRoom(roomName)
	}
	return sectorIDSum
}
