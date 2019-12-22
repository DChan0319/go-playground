package main

import "testing"

func TestHello(t *testing.T) {

	assertErrorMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q want: %q", got, want)
		}
	}

	t.Run("saying hello to myself", func(t *testing.T) {
		got := Hello("Darren")
		want := "Hello Darren, World"

		assertErrorMessage(t, got, want)
	})

	t.Run("if empty string is provided print out 'Hello, World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertErrorMessage(t, got, want)
	})
}
