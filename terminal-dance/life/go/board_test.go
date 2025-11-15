package main

import "testing"

func TestBoardNextBlinker(t *testing.T) {
	b := &Board{
		width:  5,
		height: 5,
		cells: [][]bool{
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, true, false, false},
			{false, false, true, false, false},
			{false, false, false, false, false},
		},
	}

	next := b.Next()
	expected := [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}

	for y := range expected {
		for x := range expected[y] {
			if next.cells[y][x] != expected[y][x] {
				t.Fatalf("unexpected cell state at (%d,%d)", x, y)
			}
		}
	}
}

func BenchmarkBoardNext(b *testing.B) {
	board := NewBoard(40, 20, false, 0.5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board = board.Next()
	}
}
