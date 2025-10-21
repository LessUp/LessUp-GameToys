use std::io::{self, Write};
use std::sync::{atomic::{AtomicBool, Ordering}, Arc};
use std::time::Duration;
use clap::Parser;

#[derive(Parser, Debug)]
#[command(author, version, about = "跑动小猫 (ASCII/Emoji)")]
struct Args {
    #[arg(long, default_value_t = 80usize, help = "终端宽度")]
    w: usize,
    #[arg(long, default_value_t = false, help = "使用表情风格")]
    emoji: bool,
    #[arg(long, default_value_t = 80u64, help = "每帧延迟(毫秒)")]
    speed: u64,
}

fn render_line(content: &str, width: usize, offset: usize) -> String {
    let width = if width == 0 { 80 } else { width };
    if offset >= width { return String::new(); }
    let spaces = " ".repeat(offset);
    let mut s = String::with_capacity(width);
    s.push_str(&spaces);
    let mut count = 0usize;
    for ch in content.chars() {
        if offset + count >= width { break; }
        s.push(ch);
        count += 1;
    }
    s
}

fn main() {
    let args = Args::parse();

    let running = Arc::new(AtomicBool::new(true));
    {
        let running = running.clone();
        ctrlc::set_handler(move || {
            running.store(false, Ordering::SeqCst);
        }).ok();
    }

    let ascii = ["=(^_^)=", "=(^o^)="]; 
    let emoji = ["🐈", "🐱"]; 
    let frames = if args.emoji { &emoji } else { &ascii };

    print!("\x1b[?25l");
    println!("小猫跑动中（Ctrl+C 退出）...");

    let delay = Duration::from_millis(args.speed);
    let mut idx = 0usize;
    let mut offset = 0usize;

    while running.load(Ordering::SeqCst) {
        std::thread::sleep(delay);
        print!("\r\x1b[2K{}", render_line(frames[idx % frames.len()], args.w, offset));
        io::stdout().flush().ok();
        idx = idx.wrapping_add(1);
        offset += 1;
        if offset > args.w { offset = 0; }
    }

    print!("\r\x1b[2K\x1b[?25h\n");
}
