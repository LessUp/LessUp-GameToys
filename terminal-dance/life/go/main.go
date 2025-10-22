package main

import (
    "flag"
    "fmt"
    "math/rand"
    "os"
    "os/signal"
    "syscall"
    "time"
)

type Board struct {
    width  int
    height int
    cells  [][]bool
    emoji  bool
}

func newBoard(width, height int, emoji bool, density float64) *Board {
    b := &Board{
        width:  width,
        height: height,
        cells:  make([][]bool, height),
        emoji:  emoji,
    }
    for i := range b.cells {
        b.cells[i] = make([]bool, width)
        for j := range b.cells[i] {
            if rand.Float64() < density {
                b.cells[i][j] = true
            }
        }
    }
    return b
}

func (b *Board) countNeighbors(x, y int) int {
    count := 0
    for dy := -1; dy <= 1; dy++ {
        for dx := -1; dx <= 1; dx++ {
            if dx == 0 && dy == 0 {
                continue
            }
            nx := (x + dx + b.width) % b.width
            ny := (y + dy + b.height) % b.height
            if b.cells[ny][nx] {
                count++
            }
        }
    }
    return count
}

func (b *Board) next() *Board {
    newBoard := &Board{
        width:  b.width,
        height: b.height,
        cells:  make([][]bool, b.height),
        emoji:  b.emoji,
    }
    for i := range newBoard.cells {
        newBoard.cells[i] = make([]bool, b.width)
        for j := range newBoard.cells[i] {
            neighbors := b.countNeighbors(j, i)
            alive := b.cells[i][j]
            if alive && (neighbors == 2 || neighbors == 3) {
                newBoard.cells[i][j] = true
            } else if !alive && neighbors == 3 {
                newBoard.cells[i][j] = true
            }
        }
    }
    return newBoard
}

func (b *Board) render() {
    clearScreen()
    setCursor(1, 1)

    alive := "█"
    dead := " "

    if b.emoji {
        alive = "🟢"
        dead = "⚫"
    }

    for y := 0; y < b.height; y++ {
        for x := 0; x < b.width; x++ {
            if b.cells[y][x] {
                fmt.Print(alive)
            } else {
                fmt.Print(dead)
            }
        }
        fmt.Println()
    }
}

func (b *Board) countAlive() int {
    count := 0
    for y := 0; y < b.height; y++ {
        for x := 0; x < b.width; x++ {
            if b.cells[y][x] {
                count++
            }
        }
    }
    return count
}

func clearScreen() {
    fmt.Print("\x1b[2J")
}

func setCursor(x, y int) {
    fmt.Printf("\x1b[%d;%dH", y, x)
}

func hideCursor() {
    fmt.Print("\x1b[?25l")
}

func showCursor() {
    fmt.Print("\x1b[?25h")
}

func main() {
    width := flag.Int("w", 40, "世界宽度")
    height := flag.Int("h", 20, "世界高度")
    emoji := flag.Bool("emoji", false, "使用表情风格")
    speed := flag.Int("speed", 100, "每帧延迟(毫秒)")
    density := flag.Float64("density", 0.3, "初始密度(0.0-1.0)")
    flag.Parse()

    rand.Seed(time.Now().UnixNano())

    w := *width
    if w < 4 {
        w = 4
    }
    h := *height
    if h < 4 {
        h = 4
    }
    d := *density
    if d < 0 {
        d = 0
    }
    if d > 1 {
        d = 1
    }

    board := newBoard(w, h, *emoji, d)

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    running := true
    go func() { <-c; running = false }()

    hideCursor()
    defer func() {
        clearScreen()
        showCursor()
        fmt.Println()
    }()

    delay := *speed
    if delay < 30 {
        delay = 30
    }

    ticker := time.NewTicker(time.Duration(delay) * time.Millisecond)
    defer ticker.Stop()

    generation := 0
    for running {
        board.render()
        setCursor(1, h+2)
        fmt.Printf("第 %d 代 | 存活细胞: %d | Ctrl+C 退出", generation, board.countAlive())

        <-ticker.C
        board = board.next()
        generation++
    }
}
