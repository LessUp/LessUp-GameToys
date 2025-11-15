use rand::Rng;
use std::io::{self, Write};
use std::sync::atomic::{AtomicBool, Ordering};
use std::sync::Arc;
use std::time::Duration;

pub mod config;

pub use config::Config;

pub struct Board {
    pub width: usize,
    pub height: usize,
    pub cells: Vec<bool>,
    pub emoji: bool,
}

impl Board {
    pub fn random(width: usize, height: usize, emoji: bool, density: f64) -> Self {
        let mut rng = rand::thread_rng();
        let mut cells = vec![false; width * height];
        for cell in cells.iter_mut() {
            if rng.gen::<f64>() < density {
                *cell = true;
            }
        }
        Self {
            width,
            height,
            cells,
            emoji,
        }
    }

    pub fn empty(width: usize, height: usize, emoji: bool) -> Self {
        Self {
            width,
            height,
            cells: vec![false; width * height],
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

    pub fn next(&self) -> Board {
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
                next.cells[idx] = matches!((alive, neighbors), (true, 2 | 3) | (false, 3));
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

pub fn run(cfg: Config) -> io::Result<()> {
    let cfg = cfg.sanitized();
    let mut board = Board::random(cfg.width, cfg.height, cfg.emoji, cfg.density);

    let running = Arc::new(AtomicBool::new(true));
    {
        let running = running.clone();
        ctrlc::set_handler(move || {
            running.store(false, Ordering::SeqCst);
        })
        .ok();
    }

    print!("\x1b[?25l");
    let delay = Duration::from_millis(cfg.speed_ms);
    let mut generation = 0usize;

    while running.load(Ordering::SeqCst) {
        board.render();
        println!(
            "第 {} 代 | 存活细胞: {} | Ctrl+C 退出",
            generation,
            board.count_alive()
        );
        std::thread::sleep(delay);
        board = board.next();
        generation = generation.saturating_add(1);
    }

    print!("\x1b[2J\x1b[1;1H\x1b[?25h");
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn blinker_oscillates() {
        let mut board = Board {
            width: 3,
            height: 3,
            cells: vec![false, true, false, false, true, false, false, true, false],
            emoji: false,
        };
        board = board.next();
        assert_eq!(board.cells[3..6], [true, true, true]);
    }
}
