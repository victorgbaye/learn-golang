package main

import "testing"

func TestHello (t *testing.T){
	t.Run("Saying hello to people", func(t *testing.T){
		got:= Hello("Chris")
		want:= "Hello, Chris"
	
		AssertCorrectMessage(t, got, want)
	})

	t.Run("Saying 'Hello, world' when an empty string is supplied", func(t *testing.T){
		got:= Hello("", Spanish)
		want:= "Hello, world"

		AssertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func (t *testing.T){
		got:= Hello("Hola", "Chris")
		want:= "Hola, Chris"

		AssertCorrectMessage(t, got, want)
	})

}

func AssertCorrectMessage (t testing.TB, got, want string){
	// t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}