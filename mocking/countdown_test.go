package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

const write = "write"
const sleep = "sleep"

type CoundownOperationsSpy struct {
	Calls []string
}

type Sleeper interface {
	Sleep()
}

func (s *CoundownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CoundownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CoundownOperationsSpy{})

		actual := buffer.String()
		expected := `3
2
1
Go!`

		if actual != expected {
			t.Errorf("actual: %q, expected: %q", actual, expected)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &CoundownOperationsSpy{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeperPrinter.Calls)
		}
	})
}
