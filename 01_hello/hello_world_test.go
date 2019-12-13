package main

import "testing"

func TestHelloWorld( t *testing.T)  {
	got := greeting()
	want:= "Hello, world"
	if got!=want{
		t.Errorf("got %q  wanted %q",got,want)
	}else {
		t.Log()
	}
}
