package main

import "testing"

func Test_greeter(t *testing.T) {
	got := greeting("Chris")
	want := "Hello, Chris"
	if got != want {
		t.Errorf("got %q  wanted %q", got, want)
	} else {
		t.Log()
	}
}
