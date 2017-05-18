package five

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayFivePartOne(t *testing.T) {
	passwordOne := DayFivePartOne("abc")
	assert.Equal(t, "18f47a30", passwordOne)

	passwordFinal := DayFivePartOne("wtnhxymk")
	assert.Equal(t, "2414bc77", passwordFinal)
}
