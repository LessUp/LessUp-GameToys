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

const (
	defaultFPS  = 12
	minFPS      = 1
	maxFPS      = 60
	defaultCars = 3
	minCars     = 0
	maxCars     = 20
)

var (
	asciiWheelFrames = []string{"-", "\\", "|", "/"}
	emojiWheelFrames = []string{"◐", "◓", "◑", "◒"}
	emojiSmokeFrames = []string{"", "💭", "💨", "💨"}
)

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func getTermWidth() int {
	// 优先使用环境变量；不可用时回退 80 列
	if cols := os.Getenv("COLUMNS"); cols != "" {
		if n, err := strconv.Atoi(cols); err == nil && n > 0 {
			return n
		}
	}
	return 80
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
	// 简易车头
	b.WriteString(fmt.Sprintf(" _^_o%s", w))
	// 车厢
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

func main() {
	cars := flag.Int("cars", defaultCars, "车厢数量 (0-20)")
	speed := flag.Int("speed", defaultFPS, "帧率FPS (1-60)")
	emoji := flag.Bool("emoji", false, "使用表情风格")
	flag.Parse()

	*cars = clamp(*cars, minCars, maxCars)
	*speed = clamp(*speed, minFPS, maxFPS)

	width := getTermWidth()

	// 信号处理，确保优雅恢复光标
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	hideCursor()
	defer func() {
		clearLine()
		showCursor()
		fmt.Println()
	}()

	fmt.Printf("小火车正在运行（Ctrl+C 退出） cars=%d speed=%d fps emoji=%t\n", *cars, *speed, *emoji)

	ticker := time.NewTicker(time.Second / time.Duration(*speed))
	defer ticker.Stop()

	phase := 0
	offset := 0

	for {
		select {
		case <-sig:
			return
		case <-ticker.C:
			train := buildTrain(*cars, phase, *emoji)
			clearLine()
			fmt.Print(renderLine(train, width, offset))

			phase = (phase + 1) % 1024
			offset++
			if offset > width {
				offset = 0
			}
		}
	}
}
