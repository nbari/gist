package pad_test

import (
	"github.com/nbari/gist/pad"
	"testing"
)

func BenchmarkA(b *testing.B) {
	pad.A(300)
}

func BenchmarkB(b *testing.B) {
	pad.B(300)
}

// run with:
// go test -bench=. pad_test.go
