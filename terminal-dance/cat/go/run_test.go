package main

import (
	"testing"
)

func TestRenderLine(t *testing.T) {
	got := renderLine("cat", 5, 1)
	if got != " cat" {
		t.Fatalf("unexpected render result: %q", got)
	}

	if got := renderLine("abcdef", 4, 0); got != "abcd" {
		t.Fatalf("expected truncation, got %q", got)
	}
}

func BenchmarkRenderLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = renderLine("=(^_^)=", 80, i%40)
	}
}
