package main

import "testing"

func TestDefaultGreeter(t *testing.T) {
	decisiveFunc := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q  wanted %q", got, want)
		} else {
			t.Log()
		}
	}
	t.Run("with greeter", func(t *testing.T) {
		got := greeting("Chris")
		want := "Hello, Chris"
		decisiveFunc(t,got,want)
	})
	t.Run("without greeter", func(t *testing.T) {
		got := greeting("")
		want := "Hello, World"
		decisiveFunc(t,got,want)
	})
}
