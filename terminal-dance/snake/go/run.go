package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"snake/config"
)

// Run starts the snake game loop using the provided configuration.
func Run(cfg config.Config) error {
	rand.Seed(time.Now().UnixNano())

	game := NewGame(cfg.Width, cfg.Height, cfg.Emoji)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	inputChan := make(chan rune, 10)
	go func() {
		buf := make([]byte, 1)
		for {
			if _, err := os.Stdin.Read(buf); err == nil {
				inputChan <- rune(buf[0])
			}
		}
	}()

	enableRawMode()
	hideCursor()
	defer func() {
		disableRawMode()
		showCursor()
	}()

	delay := cfg.FrameDelay
	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	game.render()

	running := true
	for running {
		select {
		case <-sig:
			running = false
		case input := <-inputChan:
			switch input {
			case 'w', 'W':
				if game.direction != Down {
					game.direction = Up
				}
			case 's', 'S':
				if game.direction != Up {
					game.direction = Down
				}
			case 'a', 'A':
				if game.direction != Right {
					game.direction = Left
				}
			case 'd', 'D':
				if game.direction != Left {
					game.direction = Right
				}
			case 'q', 'Q':
				running = false
			}
		case <-ticker.C:
			game.update()
			game.render()
			if game.gameOver {
				<-inputChan
				running = false
			}
		}
	}

	return nil
}

func clearScreen()       { fmt.Print("\x1b[2J") }
func setCursor(x, y int) { fmt.Printf("\x1b[%d;%dH", y, x) }
func hideCursor()        { fmt.Print("\x1b[?25l") }
func showCursor()        { fmt.Print("\x1b[?25h") }
func enableRawMode()     { fmt.Print("\x1b[?1049h") }
func disableRawMode()    { fmt.Print("\x1b[?1049l") }
