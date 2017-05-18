package five

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayFivePartTwo(t *testing.T) {
	passwordOne := DayFivePartTwo("abc")
	assert.Equal(t, "05ace8e3", passwordOne)

	passwordFinal := DayFivePartTwo("wtnhxymk")
	assert.Equal(t, "437e60fc", passwordFinal)
}
