package main

import "testing"

func TestHelloWithName(t *testing.T) {

	checkHello := func(got string, want string, t *testing.T) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to someone", func(t *testing.T) {
		checkHello(Hello("bob", ""), "hello bob!", t)
	})

	t.Run("when empty string say hello world!", func(t *testing.T) {
		checkHello(Hello("", ""), "hello world!", t)
	})

	t.Run("say hello to someone in spanish", func(t *testing.T) {
		checkHello(Hello("miguel", "spanish"), "hola miguel!", t)
	})

	t.Run("say hello to someone in french", func(t *testing.T) {
		checkHello(Hello("marie claire", "FrENch"), "bonjour marie claire!", t)
	})

	t.Run("when empty string say hello world! in french", func(t *testing.T) {
		checkHello(Hello("", "FRenCH"), "bonjour world!", t)
	})
}

func TestHelloNoName(t *testing.T) {
}
