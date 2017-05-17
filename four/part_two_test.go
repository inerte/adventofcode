package four

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayThreePartTwo(t *testing.T) {
	realName := GetRealName("qzmt-zixmtkozy-ivhz-343")
	assert.Equal(t, "very encrypted name", realName)

	sectorID := DayFourPartTwo(GetInput())
	assert.Equal(t, 501, sectorID)
}
