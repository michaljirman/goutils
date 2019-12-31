package safecounter

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

func TestSafeCounter(t *testing.T){
	c := SafeCounter{v : make(map[string]int)}
	for i:=0; i < 1000; i++{
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	assert.Equal(t, c.Val("somekey"), 1000, "unexpected value")
}
