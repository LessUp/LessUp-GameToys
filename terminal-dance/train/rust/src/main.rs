use clap::Parser;
use crossterm::{cursor, execute, terminal, terminal::ClearType};
use std::env;
use std::io::{stdout, Stdout, Write};
use std::sync::atomic::{AtomicBool, Ordering};
use std::sync::Arc;
use std::thread;
use std::time::Duration;
use unicode_width::UnicodeWidthChar;

const DEFAULT_FPS: u32 = 12;
const MIN_FPS: u32 = 1;
const MAX_FPS: u32 = 60;
const DEFAULT_CARS: usize = 3;
const MIN_CARS: usize = 0;
const MAX_CARS: usize = 20;
const ASCII_WHEEL_FRAMES: [&str; 4] = ["-", "\\", "|", "/"];
const EMOJI_WHEEL_FRAMES: [&str; 4] = ["◐", "◓", "◑", "◒"];
const EMOJI_SMOKE_FRAMES: [&str; 4] = ["", "💭", "💨", "💨"];

#[derive(Parser)]
#[command(author = "", version, about = "", long_about = None)]
struct Args {
    #[arg(long, default_value_t = DEFAULT_CARS)]
    cars: usize,
    #[arg(long = "speed", default_value_t = DEFAULT_FPS)]
    fps: u32,
    #[arg(long, default_value_t = false)]
    emoji: bool,
}

struct CursorGuard;

impl CursorGuard {
    fn new(stdout: &mut Stdout) -> Self {
        let _ = execute!(stdout, cursor::Hide);
        CursorGuard
    }
}

impl Drop for CursorGuard {
    fn drop(&mut self) {
        let mut stdout = stdout();
        let _ = execute!(
            stdout,
            cursor::MoveToColumn(0),
            terminal::Clear(ClearType::CurrentLine),
            cursor::Show
        );
        let _ = stdout.flush();
        println!();
    }
}

fn clamp_usize(value: usize, min: usize, max: usize) -> usize {
    value.max(min).min(max)
}

fn clamp_u32(value: u32, min: u32, max: u32) -> u32 {
    value.max(min).min(max)
}

fn get_term_width() -> usize {
    env::var("COLUMNS")
        .ok()
        .and_then(|v| v.parse::<usize>().ok())
        .filter(|&w| w > 0)
        .unwrap_or(80)
}

fn build_train(cars: usize, phase: usize, emoji: bool) -> String {
    if emoji {
        let smoke = EMOJI_SMOKE_FRAMES[phase % EMOJI_SMOKE_FRAMES.len()];
        let wheel = EMOJI_WHEEL_FRAMES[phase % EMOJI_WHEEL_FRAMES.len()];
        let mut result = String::new();
        if !smoke.is_empty() {
            result.push_str(smoke);
            result.push(' ');
        }
        result.push_str("🚂");
        result.push_str(wheel);
        for i in 0..cars {
            let idx = (phase + i + 1) % EMOJI_WHEEL_FRAMES.len();
            result.push(' ');
            result.push_str("🚃");
            result.push_str(EMOJI_WHEEL_FRAMES[idx]);
        }
        return result;
    }
    let wheel = ASCII_WHEEL_FRAMES[phase % ASCII_WHEEL_FRAMES.len()];
    let mut result = String::new();
    result.push_str(" _^_o");
    result.push_str(wheel);
    for i in 0..cars {
        let idx = (phase + i) % ASCII_WHEEL_FRAMES.len();
        result.push(' ');
        result.push_str("[=");
        result.push_str(ASCII_WHEEL_FRAMES[idx]);
        result.push_str("=]");
    }
    result
}

fn render_line(content: &str, width: usize, offset: usize) -> String {
    if width == 0 || offset >= width {
        return String::new();
    }
    let spaces = " ".repeat(offset);
    let available = width - offset;
    let mut acc = String::new();
    let mut current_width = 0;
    for ch in content.chars() {
        let w = ch.width().unwrap_or(0);
        if w == 0 {
            continue;
        }
        if current_width + w > available {
            break;
        }
        acc.push(ch);
        current_width += w;
    }
    format!("{}{}", spaces, acc)
}

fn main() {
    let mut args = Args::parse();
    args.cars = clamp_usize(args.cars, MIN_CARS, MAX_CARS);
    args.fps = clamp_u32(args.fps, MIN_FPS, MAX_FPS);
    let width = get_term_width();
    let interval = Duration::from_secs_f64(1.0 / args.fps as f64);
    let running = Arc::new(AtomicBool::new(true));
    let signal_flag = running.clone();
    let _ = ctrlc::set_handler(move || {
        signal_flag.store(false, Ordering::SeqCst);
    });
    let mut stdout = stdout();
    let _guard = CursorGuard::new(&mut stdout);
    println!(
        "小火车正在运行（Ctrl+C 退出） cars={} speed={} fps emoji={}",
        args.cars,
        args.fps,
        args.emoji
    );
    let mut phase: usize = 0;
    let mut offset: usize = 0;
    while running.load(Ordering::SeqCst) {
        let train = build_train(args.cars, phase, args.emoji);
        let line = render_line(&train, width, offset);
        let _ = execute!(
            stdout,
            cursor::MoveToColumn(0),
            terminal::Clear(ClearType::CurrentLine)
        );
        print!("{}", line);
        let _ = stdout.flush();
        phase = (phase + 1) % 1024;
        offset += 1;
        if offset > width {
            offset = 0;
        }
        thread::sleep(interval);
    }
}
