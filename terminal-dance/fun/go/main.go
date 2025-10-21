package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	asciiFrames = []string{"-", "\\", "|", "/"}
	emojiFrames = []string{"◐", "◓", "◑", "◒"}
)

func hideCursor() { fmt.Print("\x1b[?25l") }
func showCursor() { fmt.Print("\x1b[?25h") }
func clearLine()  { fmt.Print("\r\x1b[2K") }

func main() {
	emoji := flag.Bool("emoji", false, "使用表情风格")
	speed := flag.Int("speed", 80, "每帧延迟(毫秒)")
	flag.Parse()

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

	frames := asciiFrames
	if *emoji {
		frames = emojiFrames
	}

	ticker := time.NewTicker(time.Duration(*speed) * time.Millisecond)
	defer ticker.Stop()

	idx := 0
	fmt.Println("风扇旋转中（Ctrl+C 退出）...")
	for running {
		<-ticker.C
		clearLine()
		fmt.Printf("%s", frames[idx%len(frames)])
		idx++
	}
}
