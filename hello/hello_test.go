package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Andy", "")
		want := "Hello, Andy"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("saying hello to world for an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Manuel", "Spanish")
		want := "Hola, Manuel"

		assertCorrectMessage(t, got, want)
	})
	t.Run("hello world in Spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, Mundo"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Luc", "French")
		want := "Bonjour, Luc"

		assertCorrectMessage(t, got, want)
	})
	t.Run("hello world in French", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, tout le monde"

		assertCorrectMessage(t, got, want)
	})
	t.Run("should default to English for unknown language", func(t *testing.T) {
		got := Hello("", "Urdish")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
