package strings

import (
	"testing"
)

func TestRandLetterNum(t *testing.T) {
	n := 32
	a := RandLetterNum(n)
	b := RandLetterNum(n)
	c := RandLetterNum(n)

	t.Logf("%q %q %q", a, b, c)

	if len(a) != n || len(b) != n || len(c) != n {
		t.Fail()
		return
	}
	if a == b || a == c || b == c {
		t.Fail()
		return
	}
}

func BenchmarkRandLetterNum(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandLetterNum(10)
	}
}

func TestRandNum(t *testing.T) {
	n := 6
	a := RandNum(n)
	b := RandNum(n)
	c := RandNum(n)

	t.Logf("%q %q %q", a, b, c)

	if len(a) != n || len(b) != n || len(c) != n {
		t.Fail()
		return
	}
	if a == b || a == c || b == c {
		t.Fail()
		return
	}
}

func BenchmarkRandNum(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandNum(10)
	}
}
