package mocking

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(writer io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		var intAsString string = strconv.Itoa(i)
		s.Sleep()
		fmt.Fprintln(writer, intAsString)
	}

	s.Sleep()
	fmt.Fprintf(writer, "Go!")
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
