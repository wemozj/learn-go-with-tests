package main

import "testing"

func AssertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("want %s, but got %s", want, got)
	}
}

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	AssertString(t, got, want)
}