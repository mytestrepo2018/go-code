package iteration

import (
	"fmt"
	"testing"
)

func assertResult(want string, got string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestIterateChar(t *testing.T) {

	t.Run("iterate with 10 chars", func(t *testing.T) {
		want := "aaaaaaaaaa"
		got := IterateChar("a", 10)
		assertResult(want, got, t)
	})

	t.Run("iterate with 5 chars", func(t *testing.T) {
		want := "zzzzz"
		got := IterateChar("z", 5)
		assertResult(want, got, t)
	})
}

func BenchmarkIterateChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IterateChar("z", 1)
	}
}

func ExampleIterateChar() {
	iterate := IterateChar("b", 4)
	fmt.Println(iterate)
	// Output: bbbb
}
