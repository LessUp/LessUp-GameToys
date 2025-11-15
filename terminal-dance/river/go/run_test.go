package main

import "testing"

func TestClampRiverWidth(t *testing.T) {
	if got := clampRiverWidth(30, 20); got != 19 {
		t.Fatalf("expected 19, got %d", got)
	}
	if got := clampRiverWidth(5, 20); got != 5 {
		t.Fatalf("expected passthrough, got %d", got)
	}
}

func TestWaveChar(t *testing.T) {
	if c := waveChar(0); c != '~' {
		t.Fatalf("unexpected char %c", c)
	}
	if c := waveChar(1); c != '-' {
		t.Fatalf("unexpected char %c", c)
	}
	if c := waveChar(2); c != '=' {
		t.Fatalf("unexpected char %c", c)
	}
}

func BenchmarkWaveChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = waveChar(i)
	}
}
