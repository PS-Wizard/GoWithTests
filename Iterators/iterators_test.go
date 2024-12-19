package iterators

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("b", 5)
	expected := "bbbbb"
	if repeated != expected {
		t.Errorf("Expected %q got %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
