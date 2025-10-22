use clap::Parser;
use rand::Rng;
use std::io::{self, Read, Write};
use std::sync::atomic::{AtomicBool, Ordering};
use std::sync::Arc;
use std::thread;
use std::time::Duration;

#[derive(Parser, Debug)]
#[command(author, version, about = "贪吃蛇游戏 (ASCII/Emoji)")]
struct Args {
    #[arg(long, default_value_t = 20, help = "游戏宽度")]
    w: usize,
    #[arg(long, default_value_t = 10, help = "游戏高度")]
    h: usize,
    #[arg(long, default_value_t = false, help = "使用表情风格")]
    emoji: bool,
    #[arg(long, default_value_t = 150, help = "游戏速度(毫秒)")]
    speed: u64,
}

#[derive(Clone, Copy, PartialEq, Debug)]
struct Point {
    x: i32,
    y: i32,
}

#[derive(Clone, Copy, PartialEq)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

struct Game {
    width: usize,
    height: usize,
    snake: Vec<Point>,
    direction: Direction,
    food: Point,
    score: usize,
    game_over: bool,
    emoji: bool,
}

impl Game {
    fn new(width: usize, height: usize, emoji: bool) -> Self {
        let mut game = Game {
            width,
            height,
            snake: vec![Point {
                x: (width / 2) as i32,
                y: (height / 2) as i32,
            }],
            direction: Direction::Right,
            food: Point { x: 0, y: 0 },
            score: 0,
            game_over: false,
            emoji,
        };
        game.spawn_food();
        game
    }

    fn spawn_food(&mut self) {
        if self.snake.len() >= self.width * self.height {
            self.game_over = true;
            return;
        }
        let mut rng = rand::thread_rng();
        loop {
            let food = Point {
                x: rng.gen_range(0..self.width as i32),
                y: rng.gen_range(0..self.height as i32),
            };
            if !self.snake.contains(&food) {
                self.food = food;
                break;
            }
        }
    }

    fn update(&mut self) {
        if self.game_over {
            return;
        }

        let head = self.snake[0];
        let new_head = match self.direction {
            Direction::Up => Point {
                x: head.x,
                y: head.y - 1,
            },
            Direction::Down => Point {
                x: head.x,
                y: head.y + 1,
            },
            Direction::Left => Point {
                x: head.x - 1,
                y: head.y,
            },
            Direction::Right => Point {
                x: head.x + 1,
                y: head.y,
            },
        };

        if new_head.x < 0
            || new_head.x >= self.width as i32
            || new_head.y < 0
            || new_head.y >= self.height as i32
        {
            self.game_over = true;
            return;
        }

        if self.snake.contains(&new_head) {
            self.game_over = true;
            return;
        }

        self.snake.insert(0, new_head);

        if new_head == self.food {
            self.score += 1;
            self.spawn_food();
        } else {
            self.snake.pop();
        }
    }

    fn render(&self) {
        print!("\x1b[2J\x1b[1;1H");

        let (wall, snake_body, food_char, empty) = if self.emoji {
            ("🧱", "🟩", "🍎", "  ")
        } else {
            ("█", "■", "●", " ")
        };

        for _ in 0..self.width + 2 {
            print!("{}", wall);
        }
        println!();

        for y in 0..self.height {
            print!("{}", wall);
            for x in 0..self.width {
                let p = Point {
                    x: x as i32,
                    y: y as i32,
                };
                if self.snake.contains(&p) {
                    print!("{}", snake_body);
                } else if p == self.food {
                    print!("{}", food_char);
                } else {
                    print!("{}", empty);
                }
            }
            println!("{}", wall);
        }

        for _ in 0..self.width + 2 {
            print!("{}", wall);
        }
        println!();

        println!("分数: {} | 使用 WASD 控制方向，Q 退出", self.score);
        if self.game_over {
            println!("游戏结束！按任意键退出...");
        }

        io::stdout().flush().ok();
    }

    fn change_direction(&mut self, new_dir: Direction) {
        let valid = match (self.direction, new_dir) {
            (Direction::Up, Direction::Down) => false,
            (Direction::Down, Direction::Up) => false,
            (Direction::Left, Direction::Right) => false,
            (Direction::Right, Direction::Left) => false,
            _ => true,
        };
        if valid {
            self.direction = new_dir;
        }
    }
}

fn enable_raw_mode() {
    print!("\x1b[?1049h\x1b[?25l");
    io::stdout().flush().ok();
}

fn disable_raw_mode() {
    print!("\x1b[?1049l\x1b[?25h");
    io::stdout().flush().ok();
}

fn main() {
    let args = Args::parse();

    let width = args.w.max(10);
    let height = args.h.max(5);
    let mut game = Game::new(width, height, args.emoji);

    let running = Arc::new(AtomicBool::new(true));
    let running_clone = running.clone();
    ctrlc::set_handler(move || {
        running_clone.store(false, Ordering::SeqCst);
    })
    .ok();

    enable_raw_mode();

    let input_running = running.clone();
    let (tx, rx) = std::sync::mpsc::channel();
    thread::spawn(move || {
        let mut stdin = io::stdin();
        let mut buf = [0u8; 1];
        while input_running.load(Ordering::SeqCst) {
            if stdin.read(&mut buf).is_ok() {
                tx.send(buf[0] as char).ok();
            }
        }
    });

    game.render();

    let delay = Duration::from_millis(args.speed.max(30));
    let mut last_update = std::time::Instant::now();

    while running.load(Ordering::SeqCst) {
        if let Ok(ch) = rx.try_recv() {
            match ch {
                'w' | 'W' => game.change_direction(Direction::Up),
                's' | 'S' => game.change_direction(Direction::Down),
                'a' | 'A' => game.change_direction(Direction::Left),
                'd' | 'D' => game.change_direction(Direction::Right),
                'q' | 'Q' => break,
                _ => {}
            }
        }

        if last_update.elapsed() >= delay {
            game.update();
            game.render();
            last_update = std::time::Instant::now();

            if game.game_over {
                rx.recv().ok();
                break;
            }
        }

        thread::sleep(Duration::from_millis(10));
    }

    disable_raw_mode();
}
