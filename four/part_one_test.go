package four

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayThreePartOne(t *testing.T) {
	isRealRoom1 := isRealRoom("aaaaa-bbb-z-y-x-123[abxyz]")
	assert.Equal(t, true, isRealRoom1)

	isRealRoom2 := isRealRoom("a-b-c-d-e-f-g-h-987[abcde]")
	assert.Equal(t, true, isRealRoom2)

	isRealRoom3 := isRealRoom("not-a-real-room-404[oarel]")
	assert.Equal(t, true, isRealRoom3)

	isRealRoom4 := isRealRoom("totally-real-room-200[decoy]")
	assert.Equal(t, false, isRealRoom4)
}
