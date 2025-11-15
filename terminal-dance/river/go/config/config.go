package config

import (
	"flag"
	"fmt"
	"time"
)

// Config collects tunables for the river animation.
type Config struct {
	Width      int
	Height     int
	Emoji      bool
	FrameDelay time.Duration
	RiverWidth int
}

// Default returns canonical defaults.
func Default() Config {
	return Config{
		Width:      80,
		Height:     24,
		Emoji:      false,
		FrameDelay: 60 * time.Millisecond,
		RiverWidth: 8,
	}
}

// FromCLI parses CLI arguments into configuration.
func FromCLI(args []string) (Config, error) {
	cfg := Default()
	fs := flag.NewFlagSet("river", flag.ContinueOnError)
	fs.IntVar(&cfg.Width, "w", cfg.Width, "width")
	fs.IntVar(&cfg.Height, "h", cfg.Height, "height")
	fs.BoolVar(&cfg.Emoji, "emoji", cfg.Emoji, "emoji mode")
	speed := fs.Int("speed", int(cfg.FrameDelay/time.Millisecond), "frame delay ms")
	fs.IntVar(&cfg.RiverWidth, "rw", cfg.RiverWidth, "river width")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}

	if cfg.Width <= 0 || cfg.Height <= 0 {
		return cfg, fmt.Errorf("width and height must be positive")
	}
	if *speed <= 0 {
		return cfg, fmt.Errorf("speed must be positive")
	}
	if cfg.RiverWidth < 1 {
		cfg.RiverWidth = 1
	}
	cfg.FrameDelay = time.Duration(*speed) * time.Millisecond
	if cfg.Emoji && cfg.RiverWidth > 4 {
		cfg.RiverWidth = 4
	}
	return cfg, nil
}
