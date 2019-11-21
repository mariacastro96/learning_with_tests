package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloYou(t *testing.T) {
	checkTest := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Say Hello to a specific person", func(t *testing.T) {
		got := HelloYou("Maria")
		want := "Hello, Maria"

		checkTest(t, got, want)
	})

	t.Run("Say Hello to the general world", func(t *testing.T) {
		got := HelloYou("")
		want := "Hello, world"

		checkTest(t, got, want)
	})
}
