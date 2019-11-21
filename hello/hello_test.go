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
	got := HelloYou("Maria")
	want := "Hello, Maria"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
