package main

import (
	"fmt"
	"math/rand"
)

type Board struct {
	width  int
	height int
	cells  [][]bool
	emoji  bool
}

func NewBoard(width, height int, emoji bool, density float64) *Board {
	b := &Board{
		width:  width,
		height: height,
		cells:  make([][]bool, height),
		emoji:  emoji,
	}
	for i := range b.cells {
		b.cells[i] = make([]bool, width)
		for j := range b.cells[i] {
			if rand.Float64() < density {
				b.cells[i][j] = true
			}
		}
	}
	return b
}

func (b *Board) countNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx := (x + dx + b.width) % b.width
			ny := (y + dy + b.height) % b.height
			if b.cells[ny][nx] {
				count++
			}
		}
	}
	return count
}

func (b *Board) Next() *Board {
	next := &Board{
		width:  b.width,
		height: b.height,
		cells:  make([][]bool, b.height),
		emoji:  b.emoji,
	}
	for i := range next.cells {
		next.cells[i] = make([]bool, b.width)
		for j := range next.cells[i] {
			neighbors := b.countNeighbors(j, i)
			alive := b.cells[i][j]
			if alive && (neighbors == 2 || neighbors == 3) {
				next.cells[i][j] = true
			} else if !alive && neighbors == 3 {
				next.cells[i][j] = true
			}
		}
	}
	return next
}

func (b *Board) Render() {
	clearScreen()
	setCursor(1, 1)

	alive := "█"
	dead := " "

	if b.emoji {
		alive = "🟢"
		dead = "⚫"
	}

	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			if b.cells[y][x] {
				fmt.Print(alive)
			} else {
				fmt.Print(dead)
			}
		}
		fmt.Println()
	}
}

func (b *Board) CountAlive() int {
	count := 0
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			if b.cells[y][x] {
				count++
			}
		}
	}
	return count
}
