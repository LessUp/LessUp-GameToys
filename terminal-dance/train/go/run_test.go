package main

import "testing"

func TestBuildTrainASCII(t *testing.T) {
	train := buildTrain(2, 0, false)
	if train == "" || train[0] != ' ' {
		t.Fatalf("unexpected ascii train: %q", train)
	}
}

func TestBuildTrainEmoji(t *testing.T) {
	train := buildTrain(1, 1, true)
	if len([]rune(train)) < 4 {
		t.Fatalf("unexpected emoji train: %q", train)
	}
}

func BenchmarkBuildTrain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = buildTrain(5, i, false)
	}
}
