package two

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayTwo(t *testing.T) {
	codeOne := DayTwo(`ULL
RRDDD
LURDL
UUUUD`)
	assert.Equal(t, codeOne, [4]float64{1, 9, 8, 5})
}
