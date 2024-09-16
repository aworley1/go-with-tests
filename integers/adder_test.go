package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("2 + 2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertCorrectResult(t, sum, expected)
	})
	t.Run("3 + 4", func(t *testing.T) {
		sum := Add(3, 4)
		expected := 7

		assertCorrectResult(t, sum, expected)
	})
}

func assertCorrectResult(t *testing.T, sum int, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
