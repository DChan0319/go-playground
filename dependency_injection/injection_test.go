package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Darren")

	actual := buffer.String()
	expected := "Hello, Darren"

	if actual != expected {
		t.Errorf("actual: %q, expected: %q", actual, expected)
	}
}
