package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	DefaultFPS  = 12
	MinFPS      = 1
	MaxFPS      = 60
	DefaultCars = 3
	MinCars     = 0
	MaxCars     = 20
)

// Config describes tunables for the train animation.
type Config struct {
	Cars     int
	FPS      int
	Emoji    bool
	Width    int
	Interval time.Duration
}

// Default returns defaults using the current terminal width.
func Default() Config {
	return Config{
		Cars:     DefaultCars,
		FPS:      DefaultFPS,
		Emoji:    false,
		Width:    termWidth(),
		Interval: time.Second / DefaultFPS,
	}
}

// FromCLI parses CLI arguments into a configuration instance.
func FromCLI(args []string) (Config, error) {
	cfg := Default()
	fs := flag.NewFlagSet("train", flag.ContinueOnError)
	fs.IntVar(&cfg.Cars, "cars", cfg.Cars, "车厢数量 (0-20)")
	fs.IntVar(&cfg.FPS, "speed", cfg.FPS, "帧率FPS (1-60)")
	fs.BoolVar(&cfg.Emoji, "emoji", cfg.Emoji, "使用表情风格")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}

	if cfg.Cars < MinCars || cfg.Cars > MaxCars {
		return cfg, fmt.Errorf("cars must be between %d and %d", MinCars, MaxCars)
	}
	if cfg.FPS < MinFPS || cfg.FPS > MaxFPS {
		return cfg, fmt.Errorf("speed must be between %d and %d", MinFPS, MaxFPS)
	}
	cfg.Interval = time.Second / time.Duration(cfg.FPS)
	cfg.Width = termWidth()
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
