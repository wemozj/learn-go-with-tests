package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("wemo")
	want := "Hello, wemo"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}