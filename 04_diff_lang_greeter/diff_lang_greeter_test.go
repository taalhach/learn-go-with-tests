package main

import "testing"

func TestGreeter(t *testing.T)  {
	got := greeter("Chris","")
	want := "Hello, Chris"
	if got != want {
		t.Errorf("got %q  wanted %q", got, want)
	} else {
		t.Log()
	}
}

