package main

import "testing"

func TestGameUpdateGrowth(t *testing.T) {
	g := NewGame(10, 5, false)
	g.food = Point{g.snake[0].x + 1, g.snake[0].y}
	g.update()
	if len(g.snake) != 2 {
		t.Fatalf("expected snake to grow, length=%d", len(g.snake))
	}
	if g.score != 1 {
		t.Fatalf("expected score increment, got %d", g.score)
	}
}

func TestGameSelfCollision(t *testing.T) {
	g := NewGame(5, 5, false)
	g.snake = []Point{{2, 2}, {1, 2}, {1, 1}, {2, 1}}
	g.direction = Up
	g.update()
	if !g.gameOver {
		t.Fatalf("expected game over on self collision")
	}
}

func BenchmarkGameUpdate(b *testing.B) {
	g := NewGame(20, 10, false)
	for i := 0; i < b.N; i++ {
		g.update()
	}
}
