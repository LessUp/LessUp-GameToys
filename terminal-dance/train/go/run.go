package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"train/config"
)

var (
	asciiWheelFrames = []string{"-", "\\", "|", "/"}
	emojiWheelFrames = []string{"◐", "◓", "◑", "◒"}
	emojiSmokeFrames = []string{"", "💭", "💨", "💨"}
)

// Run animates the terminal train using the provided configuration.
func Run(cfg config.Config) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	hideCursor()
	defer func() {
		clearLine()
		showCursor()
		fmt.Println()
	}()

	fmt.Printf("小火车正在运行（Ctrl+C 退出） cars=%d speed=%d fps emoji=%t\n", cfg.Cars, cfg.FPS, cfg.Emoji)

	ticker := time.NewTicker(cfg.Interval)
	defer ticker.Stop()

	phase := 0
	offset := 0

	for {
		select {
		case <-sig:
			return nil
		case <-ticker.C:
			train := buildTrain(cfg.Cars, phase, cfg.Emoji)
			clearLine()
			fmt.Print(renderLine(train, cfg.Width, offset))

			phase = (phase + 1) % 1024
			offset++
			if offset > cfg.Width {
				offset = 0
			}
		}
	}
}

func buildTrain(cars, phase int, emoji bool) string {
	if emoji {
		smoke := emojiSmokeFrames[phase%len(emojiSmokeFrames)]
		wheel := emojiWheelFrames[phase%len(emojiWheelFrames)]
		var b strings.Builder
		if smoke != "" {
			b.WriteString(smoke)
			b.WriteString(" ")
		}
		b.WriteString("🚂")
		b.WriteString(wheel)
		for i := 0; i < cars; i++ {
			wi := emojiWheelFrames[(phase+i+1)%len(emojiWheelFrames)]
			b.WriteString(" ")
			b.WriteString("🚃")
			b.WriteString(wi)
		}
		return b.String()
	}

	w := asciiWheelFrames[phase%len(asciiWheelFrames)]
	var b strings.Builder
	b.WriteString(fmt.Sprintf(" _^_o%s", w))
	for i := 0; i < cars; i++ {
		wi := asciiWheelFrames[(phase+i)%len(asciiWheelFrames)]
		b.WriteString(" ")
		b.WriteString(fmt.Sprintf("[=%s=]", wi))
	}
	return b.String()
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
