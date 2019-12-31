package sum

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestConcurrentSum(t *testing.T) {
	got := ConcurrentSum([]int{1, 2, 3, 4, 5})
	assert.Equal(t, got, 15, "unexpected sum")
}