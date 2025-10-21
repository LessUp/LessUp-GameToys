package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	w := flag.Int("w", 80, "width")
	h := flag.Int("h", 24, "height")
	emoji := flag.Bool("emoji", false, "emoji mode")
	speed := flag.Int("speed", 60, "frame delay ms")
	rw := flag.Int("rw", 8, "river width")
	flag.Parse()

	riverW := *rw
	if *emoji && riverW > 4 {
		riverW = 4
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	running := true
	go func() { <-c; running = false }()

	fmt.Print("\x1b[?25l")
	defer fmt.Print("\x1b[?25h\x1b[0m\n")

	freq := 0.20
	ticker := time.NewTicker(time.Duration(*speed) * time.Millisecond)
	defer ticker.Stop()

	frame := 0
	var b strings.Builder
	for running {
		b.Reset()
		b.WriteString("\x1b[H\x1b[2J")
		width := *w
		height := *h
		riverWidth := riverW
        if riverWidth >= height {
            riverWidth = height - 1
        }
		halfH := riverWidth / 2
		amp := int(float64(height - riverWidth - 1) * 0.25)
		if amp < 1 {
			amp = 1
		}
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				f := float64(col) * freq
				s := math.Sin(f)
				center := height / 2 + int(float64(amp) * s)
				inRiver := row >= (center - halfH) && row < (center + halfH)
				if inRiver {
					if *emoji {
						if (col + frame) % 2 == 0 {
							b.WriteString("🌊")
						} else {
							b.WriteString("💧")
						}
					} else {
						val := col - frame
						idx := ((val % 3) + 3) % 3
						switch idx {
						case 0:
							b.WriteByte('~')
						case 1:
							b.WriteByte('-')
						default:
							b.WriteByte('=')
						}
					}
				} else {
					b.WriteByte(' ')
				}
			}
			b.WriteByte('\n')
		}
		fmt.Print(b.String())
		frame++
		<-ticker.C
	}
}
