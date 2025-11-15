package config

import (
	"flag"
	"fmt"
	"time"
)

// Config groups tunables for the Conway's Game of Life animation.
type Config struct {
	Width      int
	Height     int
	Emoji      bool
	FrameDelay time.Duration
	Density    float64
}

// Default returns the canonical defaults.
func Default() Config {
	return Config{
		Width:      40,
		Height:     20,
		Emoji:      false,
		FrameDelay: 100 * time.Millisecond,
		Density:    0.3,
	}
}

// FromCLI parses CLI arguments into a configuration struct.
func FromCLI(args []string) (Config, error) {
	cfg := Default()
	fs := flag.NewFlagSet("life", flag.ContinueOnError)
	fs.IntVar(&cfg.Width, "w", cfg.Width, "世界宽度")
	fs.IntVar(&cfg.Height, "h", cfg.Height, "世界高度")
	fs.BoolVar(&cfg.Emoji, "emoji", cfg.Emoji, "使用表情风格")
	speed := fs.Int("speed", int(cfg.FrameDelay/time.Millisecond), "每帧延迟(毫秒)")
	fs.Float64Var(&cfg.Density, "density", cfg.Density, "初始密度(0.0-1.0)")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}

	if cfg.Width < 4 {
		cfg.Width = 4
	}
	if cfg.Height < 4 {
		cfg.Height = 4
	}
	if cfg.Density < 0 || cfg.Density > 1 {
		return cfg, fmt.Errorf("density must be between 0 and 1")
	}
	if *speed < 30 {
		*speed = 30
	}
	cfg.FrameDelay = time.Duration(*speed) * time.Millisecond
	return cfg, nil
}
