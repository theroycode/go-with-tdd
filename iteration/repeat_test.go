package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 7)
	want := "aaaaaaa"

	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func ExampleRepeat() {
	got := Repeat("b", 6)
	fmt.Print(got)
	// output: bbbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}
