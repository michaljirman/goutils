package add_numbers_new

import "testing"

func TestAddNumbers(t *testing.T){
	t.Log("started")
	AddNumbers()
	if len(numbers) != 20 {
		t.Fail()
	}
}