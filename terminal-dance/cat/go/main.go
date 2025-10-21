package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	asciiCats = []string{"=(^_^)=", "=(^o^)="}
	emojiCats = []string{"🐈", "🐱"}
)

func getTermWidth() int {
	if cols := os.Getenv("COLUMNS"); cols != "" {
		if n, err := strconv.Atoi(cols); err == nil && n > 0 {
			return n
		}
	}
	return 80
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

func main() {
	w := flag.Int("w", getTermWidth(), "终端宽度")
	emoji := flag.Bool("emoji", false, "使用表情风格")
	speed := flag.Int("speed", 80, "每帧延迟(毫秒)")
	flag.Parse()

	frames := asciiCats
	if *emoji {
		frames = emojiCats
	}

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

	ticker := time.NewTicker(time.Duration(*speed) * time.Millisecond)
	defer ticker.Stop()

	idx := 0
	offset := 0
	fmt.Println("小猫跑动中（Ctrl+C 退出）...")
	for running {
		<-ticker.C
		clearLine()
		cat := frames[idx%len(frames)]
		fmt.Print(renderLine(cat, *w, offset))
		idx++
		offset++
		if offset > *w {
			offset = 0
		}
	}
}
