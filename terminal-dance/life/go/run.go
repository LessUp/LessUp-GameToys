package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"life/config"
)

// Run executes the Game of Life animation using the provided configuration.
func Run(cfg config.Config) error {
	rand.Seed(time.Now().UnixNano())
	board := NewBoard(cfg.Width, cfg.Height, cfg.Emoji, cfg.Density)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	running := true
	go func() { <-c; running = false }()

	hideCursor()
	defer func() {
		clearScreen()
		showCursor()
		fmt.Println()
	}()

	ticker := time.NewTicker(cfg.FrameDelay)
	defer ticker.Stop()

	generation := 0
	for running {
		board.Render()
		setCursor(1, cfg.Height+2)
		fmt.Printf("第 %d 代 | 存活细胞: %d | Ctrl+C 退出", generation, board.CountAlive())

		<-ticker.C
		board = board.Next()
		generation++
	}

	return nil
}

func clearScreen()       { fmt.Print("\x1b[2J") }
func setCursor(x, y int) { fmt.Printf("\x1b[%d;%dH", y, x) }
func hideCursor()        { fmt.Print("\x1b[?25l") }
func showCursor()        { fmt.Print("\x1b[?25h") }
