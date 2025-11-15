package main

import (
	"testing"

	"fun/config"
)

func TestFrameSet(t *testing.T) {
	if got := frameSet(false); len(got) != 4 || got[0] != "-" {
		t.Fatalf("unexpected ascii frame set: %v", got)
	}

	if got := frameSet(true); len(got) != 4 || got[0] != "◐" {
		t.Fatalf("unexpected emoji frame set: %v", got)
	}
}

func BenchmarkFrameLoop(b *testing.B) {
	cfg := config.Config{Emoji: false, FrameDelay: 1}
	frames := frameSet(cfg.Emoji)
	idx := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		idx = (idx + 1) % len(frames)
	}
	_ = idx
}
