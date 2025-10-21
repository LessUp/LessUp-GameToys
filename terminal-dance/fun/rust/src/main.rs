use std::io::{self, Write};
use std::sync::{atomic::{AtomicBool, Ordering}, Arc};
use std::time::Duration;
use clap::Parser;

#[derive(Parser, Debug)]
#[command(author, version, about = "旋转风扇 (ASCII/Emoji)")]
struct Args {
    #[arg(long, default_value_t = false, help = "使用表情风格")]
    emoji: bool,
    #[arg(long, default_value_t = 80u64, help = "每帧延迟(毫秒)")]
    speed: u64,
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

    let ascii = ["-", "\\", "|", "/"]; 
    let emoji = ["◐", "◓", "◑", "◒"]; 
    let frames = if args.emoji { &emoji } else { &ascii };

    print!("\x1b[?25l");
    println!("风扇旋转中（Ctrl+C 退出）...");

    let mut idx = 0usize;
    let delay = Duration::from_millis(args.speed);

    while running.load(Ordering::SeqCst) {
        std::thread::sleep(delay);
        print!("\r\x1b[2K{}", frames[idx % frames.len()]);
        io::stdout().flush().ok();
        idx = idx.wrapping_add(1);
    }

    print!("\r\x1b[2K\x1b[?25h\n");
}
