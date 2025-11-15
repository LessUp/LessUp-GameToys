package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fun/config"
)

var (
	asciiFrames = []string{"-", "\\", "|", "/"}
	emojiFrames = []string{"◐", "◓", "◑", "◒"}
)

// Run executes the fan animation using the provided configuration.
func Run(cfg config.Config) error {
	frames := frameSet(cfg.Emoji)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	running := true
	go func() { <-c; running = false }()

	hideCursor()
	defer func() {
		clearLine()
		showCursor()
		fmt.Println()
	}()

	ticker := time.NewTicker(cfg.FrameDelay)
	defer ticker.Stop()

	idx := 0
	fmt.Println("风扇旋转中（Ctrl+C 退出）...")
	for running {
		<-ticker.C
		clearLine()
		fmt.Printf("%s", frames[idx%len(frames)])
		idx++
	}

	return nil
}

func hideCursor() { fmt.Print("\x1b[?25l") }
func showCursor() { fmt.Print("\x1b[?25h") }
func clearLine()  { fmt.Print("\r\x1b[2K") }

func frameSet(emoji bool) []string {
	if emoji {
		return emojiFrames
	}
	return asciiFrames
}
