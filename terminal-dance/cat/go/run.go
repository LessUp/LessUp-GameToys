package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"cat/config"
)

var (
	asciiCats = []string{"=(^_^)=", "=(^o^)="}
	emojiCats = []string{"🐈", "🐱"}
)

// Run launches the cat animation using the provided configuration.
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
	offset := 0
	fmt.Println("小猫跑动中（Ctrl+C 退出）...")
	for running {
		<-ticker.C
		clearLine()
		cat := frames[idx%len(frames)]
		fmt.Print(renderLine(cat, cfg.Width, offset))
		idx++
		offset++
		if offset > cfg.Width {
			offset = 0
		}
	}

	return nil
}

func frameSet(emoji bool) []string {
	if emoji {
		return emojiCats
	}
	return asciiCats
}

func renderLine(content string, width, offset int) string {
	if width <= 0 {
		width = 80
	}
	if offset < 0 {
		offset = 0
	}
	if offset >= width {
		return ""
	}
	spaces := strings.Repeat(" ", offset)
	available := width - offset
	runes := []rune(content)
	if len(runes) > available {
		runes = runes[:available]
	}
	return spaces + string(runes)
}

func hideCursor() { fmt.Print("\x1b[?25l") }
func showCursor() { fmt.Print("\x1b[?25h") }
func clearLine()  { fmt.Print("\r\x1b[2K") }
