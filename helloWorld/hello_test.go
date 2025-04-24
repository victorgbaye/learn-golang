package main

import "testing"

func TestHello (t *testing.T){
	t.Run("Saying hello world to people", func(t *testing.T)  {
		got:= Hello("Chris", "English")
		want:="Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say hello world when an empty string is supplied", func(t *testing.T){
		got:= Hello("", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}