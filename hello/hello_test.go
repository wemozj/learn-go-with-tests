package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // 测试失败时报告所在的行，而不是辅助函数内部的行
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish: ", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Chinese: ", func(t *testing.T) {
		got := Hello("世界", "Chinese")
		want := "你好, 世界"
		assertCorrectMessage(t, got, want)
	})

}