package config

import (
	"flag"
	"time"
)

// Config describes tunables for the snake game.
type Config struct {
	Width      int
	Height     int
	Emoji      bool
	FrameDelay time.Duration
}

// Default returns baseline values.
func Default() Config {
	return Config{
		Width:      20,
		Height:     10,
		Emoji:      false,
		FrameDelay: 150 * time.Millisecond,
	}
}

// FromCLI parses CLI arguments into a configuration instance.
func FromCLI(args []string) (Config, error) {
	cfg := Default()
	fs := flag.NewFlagSet("snake", flag.ContinueOnError)
	fs.IntVar(&cfg.Width, "w", cfg.Width, "游戏宽度")
	fs.IntVar(&cfg.Height, "h", cfg.Height, "游戏高度")
	fs.BoolVar(&cfg.Emoji, "emoji", cfg.Emoji, "使用表情风格")
	speed := fs.Int("speed", int(cfg.FrameDelay/time.Millisecond), "游戏速度(毫秒)")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}

	if cfg.Width < 10 {
		cfg.Width = 10
	}
	if cfg.Height < 5 {
		cfg.Height = 5
	}
	if *speed < 30 {
		*speed = 30
	}
	cfg.FrameDelay = time.Duration(*speed) * time.Millisecond
	return cfg, nil
}
