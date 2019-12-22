package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Darren")
	want := "Hello Darren, World"

	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}
