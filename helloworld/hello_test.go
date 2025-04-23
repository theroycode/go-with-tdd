package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Rohan", "")
		want := "Hello Rohan"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Rohan", "Spanish")
		want := "Hola Rohan"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("Rohan", "French")
		want := "Bonjour Rohan"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
