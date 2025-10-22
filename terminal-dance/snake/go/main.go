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

type Point struct {
    x, y int
}

type Direction int

const (
    Up Direction = iota
    Down
    Left
    Right
)

type Game struct {
    width     int
    height    int
    snake     []Point
    direction Direction
    food      Point
    score     int
    gameOver  bool
    emoji     bool
}

func (g *Game) init() {
    g.snake = []Point{{g.width / 2, g.height / 2}}
    g.direction = Right
    g.spawnFood()
    g.score = 0
    g.gameOver = false
}

func (g *Game) spawnFood() {
    if len(g.snake) >= g.width*g.height {
        g.gameOver = true
        return
    }
    for {
        g.food = Point{rand.Intn(g.width), rand.Intn(g.height)}
        collision := false
        for _, s := range g.snake {
            if s == g.food {
                collision = true
                break
            }
        }
        if !collision {
            break
        }
    }
}

func (g *Game) update() {
    if g.gameOver {
        return
    }

    head := g.snake[0]
    var newHead Point

    switch g.direction {
    case Up:
        newHead = Point{head.x, head.y - 1}
    case Down:
        newHead = Point{head.x, head.y + 1}
    case Left:
        newHead = Point{head.x - 1, head.y}
    case Right:
        newHead = Point{head.x + 1, head.y}
    }

    if newHead.x < 0 || newHead.x >= g.width || newHead.y < 0 || newHead.y >= g.height {
        g.gameOver = true
        return
    }

    for _, s := range g.snake {
        if s == newHead {
            g.gameOver = true
            return
        }
    }

    g.snake = append([]Point{newHead}, g.snake...)

    if newHead == g.food {
        g.score++
        g.spawnFood()
    } else {
        g.snake = g.snake[:len(g.snake)-1]
    }
}

func (g *Game) render() {
    clearScreen()
    setCursor(1, 1)

    wall := "█"
    snakeBody := "■"
    foodChar := "●"
    empty := " "

    if g.emoji {
        wall = "🧱"
        snakeBody = "🟩"
        foodChar = "🍎"
        empty = "  "
    }

    for x := 0; x < g.width+2; x++ {
        fmt.Print(wall)
    }
    fmt.Println()

    for y := 0; y < g.height; y++ {
        fmt.Print(wall)
        for x := 0; x < g.width; x++ {
            p := Point{x, y}
            isSnake := false
            for _, s := range g.snake {
                if s == p {
                    isSnake = true
                    break
                }
            }
            if isSnake {
                fmt.Print(snakeBody)
            } else if p == g.food {
                fmt.Print(foodChar)
            } else {
                fmt.Print(empty)
            }
        }
        fmt.Println(wall)
    }

    for x := 0; x < g.width+2; x++ {
        fmt.Print(wall)
    }
    fmt.Println()

    fmt.Printf("分数: %d | 使用 WASD 控制方向，Q 退出\n", g.score)
    if g.gameOver {
        fmt.Println("游戏结束！按任意键退出...")
    }
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

func enableRawMode() {
    fmt.Print("\x1b[?1049h")
}

func disableRawMode() {
    fmt.Print("\x1b[?1049l")
}

func main() {
    width := flag.Int("w", 20, "游戏宽度")
    height := flag.Int("h", 10, "游戏高度")
    emoji := flag.Bool("emoji", false, "使用表情风格")
    speed := flag.Int("speed", 150, "游戏速度(毫秒)")
    flag.Parse()

    rand.Seed(time.Now().UnixNano())

    w := *width
    if w < 10 {
        w = 10
    }
    h := *height
    if h < 5 {
        h = 5
    }

    game := &Game{
        width:  w,
        height: h,
        emoji:  *emoji,
    }
    game.init()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)

    inputChan := make(chan rune, 10)
    go func() {
        buf := make([]byte, 1)
        for {
            os.Stdin.Read(buf)
            inputChan <- rune(buf[0])
        }
    }()

    enableRawMode()
    hideCursor()
    defer func() {
        disableRawMode()
        showCursor()
    }()

    delay := *speed
    if delay < 30 {
        delay = 30
    }

    ticker := time.NewTicker(time.Duration(delay) * time.Millisecond)
    defer ticker.Stop()

    game.render()

    running := true
    for running {
        select {
        case <-c:
            running = false
        case input := <-inputChan:
            switch input {
            case 'w', 'W':
                if game.direction != Down {
                    game.direction = Up
                }
            case 's', 'S':
                if game.direction != Up {
                    game.direction = Down
                }
            case 'a', 'A':
                if game.direction != Right {
                    game.direction = Left
                }
            case 'd', 'D':
                if game.direction != Left {
                    game.direction = Right
                }
            case 'q', 'Q':
                running = false
            }
        case <-ticker.C:
            game.update()
            game.render()
            if game.gameOver {
                <-inputChan
                running = false
            }
        }
    }
}
