package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "Dave")

	got := buf.String()
	want := "Hello, Dave"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
