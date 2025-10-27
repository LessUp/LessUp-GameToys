use clap::Parser;
use rand::Rng;
use std::io::{self, Write};
use std::sync::atomic::{AtomicBool, Ordering};
use std::sync::Arc;
use std::time::Duration;

#[derive(Parser, Debug)]
#[command(author, version, about = "康威生命游戏 (ASCII/Emoji)")]
struct Args {
    #[arg(long, default_value_t = 40usize, help = "世界宽度")]
    w: usize,
    #[arg(long, default_value_t = 20usize, help = "世界高度")]
    h: usize,
    #[arg(long, default_value_t = false, help = "使用表情风格")]
    emoji: bool,
    #[arg(long, default_value_t = 100u64, help = "每帧延迟(毫秒)")]
    speed: u64,
    #[arg(long, default_value_t = 0.3f64, help = "初始密度(0.0-1.0)")]
    density: f64,
}

struct Board {
    width: usize,
    height: usize,
    cells: Vec<bool>,
    emoji: bool,
}

impl Board {
    fn new(width: usize, height: usize, emoji: bool, density: f64) -> Self {
        let mut rng = rand::thread_rng();
        let mut cells = vec![false; width * height];
        for cell in cells.iter_mut() {
            if rng.gen::<f64>() < density {
                *cell = true;
            }
        }
        Board {
            width,
            height,
            cells,
            emoji,
        }
    }

    fn index(&self, x: isize, y: isize) -> usize {
        let w = self.width as isize;
        let h = self.height as isize;
        let nx = ((x % w) + w) % w;
        let ny = ((y % h) + h) % h;
        (ny as usize) * self.width + nx as usize
    }

    fn is_alive(&self, x: isize, y: isize) -> bool {
        self.cells[self.index(x, y)]
    }

    fn neighbor_count(&self, x: isize, y: isize) -> usize {
        let mut count = 0;
        for dy in -1..=1 {
            for dx in -1..=1 {
                if dx == 0 && dy == 0 {
                    continue;
                }
                if self.is_alive(x + dx, y + dy) {
                    count += 1;
                }
            }
        }
        count
    }

    fn next(&self) -> Board {
        let mut next = Board {
            width: self.width,
            height: self.height,
            cells: vec![false; self.width * self.height],
            emoji: self.emoji,
        };

        for y in 0..self.height {
            for x in 0..self.width {
                let idx = y * self.width + x;
                let neighbors = self.neighbor_count(x as isize, y as isize);
                let alive = self.cells[idx];
                next.cells[idx] = match (alive, neighbors) {
                    (true, 2 | 3) => true,
                    (false, 3) => true,
                    _ => false,
                };
            }
        }
        next
    }

    fn render(&self) {
        print!("\x1b[2J\x1b[1;1H");
        let (alive, dead) = if self.emoji {
            ("🟢", "⚫")
        } else {
            ("█", " ")
        };

        for y in 0..self.height {
            for x in 0..self.width {
                if self.cells[y * self.width + x] {
                    print!("{}", alive);
                } else {
                    print!("{}", dead);
                }
            }
            println!();
        }
        io::stdout().flush().ok();
    }

    fn count_alive(&self) -> usize {
        self.cells.iter().filter(|&&c| c).count()
    }
}

fn hide_cursor() {
    print!("\x1b[?25l");
}

fn show_cursor() {
    print!("\x1b[?25h");
}

fn main() {
    let args = Args::parse();
    
    let width = args.w.max(4);
    let height = args.h.max(4);
    let density = args.density.clamp(0.0, 1.0);
    
    let mut board = Board::new(width, height, args.emoji, density);

    let running = Arc::new(AtomicBool::new(true));
    {
        let running = running.clone();
        ctrlc::set_handler(move || {
            running.store(false, Ordering::SeqCst);
        })
        .ok();
    }

    hide_cursor();
    let delay = Duration::from_millis(args.speed.max(30));
    let mut generation = 0usize;

    while running.load(Ordering::SeqCst) {
        board.render();
        println!("第 {} 代 | 存活细胞: {} | Ctrl+C 退出", generation, board.count_alive());
        std::thread::sleep(delay);
        board = board.next();
        generation = generation.saturating_add(1);
    }

    print!("\x1b[2J\x1b[1;1H");
    show_cursor();
}
