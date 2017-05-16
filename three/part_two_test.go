package three

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayThreePartTwo(t *testing.T) {
	trianglesQty := DayThreePartTwo(GetInput())
	assert.Equal(t, 1921, trianglesQty)
}
