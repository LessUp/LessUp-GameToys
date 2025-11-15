package config

import (
	"flag"
	"fmt"
	"time"
)

// Config collects the tunable parameters for the fan animation.
type Config struct {
	Emoji      bool
	FrameDelay time.Duration
}

// Default returns the default configuration for the animation.
func Default() Config {
	return Config{
		Emoji:      false,
		FrameDelay: 80 * time.Millisecond,
	}
}

// FromCLI parses CLI arguments into a Config instance.
func FromCLI(args []string) (Config, error) {
	cfg := Default()
	fs := flag.NewFlagSet("fun", flag.ContinueOnError)
	fs.BoolVar(&cfg.Emoji, "emoji", cfg.Emoji, "使用表情风格")
	speed := fs.Int("speed", int(cfg.FrameDelay/time.Millisecond), "每帧延迟(毫秒)")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}

	if *speed <= 0 {
		return cfg, fmt.Errorf("speed must be greater than 0")
	}
	cfg.FrameDelay = time.Duration(*speed) * time.Millisecond

	return cfg, nil
}
