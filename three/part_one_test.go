package three

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayThreePartOne(t *testing.T) {
	lineIsTriangle := LineIsTriangle("5 10 25")
	assert.Equal(t, false, lineIsTriangle)

	trianglesQty := DayThreePartOne(GetInput())
	assert.Equal(t, 1050, trianglesQty)
}
