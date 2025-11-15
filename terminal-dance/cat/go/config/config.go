package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config stores CLI tunables for the running cat animation.
type Config struct {
	Width      int
	Emoji      bool
	FrameDelay time.Duration
}

// Default returns the default configuration based on the environment.
func Default() Config {
	return Config{
		Width:      termWidth(),
		Emoji:      false,
		FrameDelay: 80 * time.Millisecond,
	}
}

// FromCLI parses command line arguments into a Config instance.
func FromCLI(args []string) (Config, error) {
	cfg := Default()
	fs := flag.NewFlagSet("cat", flag.ContinueOnError)
	fs.IntVar(&cfg.Width, "w", cfg.Width, "终端宽度")
	fs.BoolVar(&cfg.Emoji, "emoji", cfg.Emoji, "使用表情风格")
	speed := fs.Int("speed", int(cfg.FrameDelay/time.Millisecond), "每帧延迟(毫秒)")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}

	if cfg.Width <= 0 {
		return cfg, fmt.Errorf("width must be greater than 0")
	}
	if *speed <= 0 {
		return cfg, fmt.Errorf("speed must be greater than 0")
	}
	cfg.FrameDelay = time.Duration(*speed) * time.Millisecond
	return cfg, nil
}

func termWidth() int {
	if cols := os.Getenv("COLUMNS"); cols != "" {
		if n, err := strconv.Atoi(cols); err == nil && n > 0 {
			return n
		}
	}
	return 80
}
