//following tutorial on https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world
package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("hello with non-empty argument", func(t *testing.T) {
		got := Hello("Abiola")
		want := "Hello, Abiola"

		assertCorrectMessage(t, got, want)
	})

	t.Run("hello with empty argument", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Testing"

		assertCorrectMessage(t, got, want)
	})

}
