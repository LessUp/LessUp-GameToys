package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"river/config"
)

const waveFrequency = 0.20

// Run animates the terminal river using the provided configuration.
func Run(cfg config.Config) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	running := true
	go func() { <-c; running = false }()

	fmt.Print("\x1b[?25l")
	defer fmt.Print("\x1b[?25h\x1b[0m\n")

	ticker := time.NewTicker(cfg.FrameDelay)
	defer ticker.Stop()

	frame := 0
	var b strings.Builder
	for running {
		b.Reset()
		b.WriteString("\x1b[H\x1b[2J")
		width := cfg.Width
		height := cfg.Height
		riverWidth := clampRiverWidth(cfg.RiverWidth, height)
		halfH := riverWidth / 2
		amp := int(float64(height-riverWidth-1) * 0.25)
		if amp < 1 {
			amp = 1
		}
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				center := waveCenter(height, riverWidth, amp, frame, col)
				if row >= (center-halfH) && row < (center+halfH) {
					if cfg.Emoji {
						if (col+frame)%2 == 0 {
							b.WriteString("🌊")
						} else {
							b.WriteString("💧")
						}
					} else {
						b.WriteByte(waveChar(col - frame))
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

	return nil
}

func clampRiverWidth(riverWidth, height int) int {
	if riverWidth >= height {
		return height - 1
	}
	return riverWidth
}

func waveCenter(height, riverWidth, amp, frame, col int) int {
	f := float64(col) * waveFrequency
	s := math.Sin(f)
	return height/2 + int(float64(amp)*s)
}

func waveChar(val int) byte {
	idx := ((val % 3) + 3) % 3
	switch idx {
	case 0:
		return '~'
	case 1:
		return '-'
	default:
		return '='
	}
}
