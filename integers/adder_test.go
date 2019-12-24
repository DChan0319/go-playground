package integers

import "testing"

func TestAddr(t *testing.T) {
	sum := Add(2, 4)
	expected := 6

	if sum != expected {
		t.Errorf("Expected: %d Actual: %d", expected, sum)
	}
}
