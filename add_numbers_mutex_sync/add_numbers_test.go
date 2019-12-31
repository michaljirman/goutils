package add_numbers_mutex_sync

import "testing"

func TestAddNumbers(t *testing.T) {
	AddNumbers()
	if len(numbers) != 10 {
		t.Fail()
	}
}
