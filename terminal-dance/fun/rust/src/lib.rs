use std::io::{self, Write};
use std::sync::{
    atomic::{AtomicBool, Ordering},
    Arc,
};
use std::time::Duration;

pub mod config;

pub use config::Config;

const ASCII_FRAMES: [&str; 4] = ["-", "\\", "|", "/"];
const EMOJI_FRAMES: [&str; 4] = ["◐", "◓", "◑", "◒"];

pub fn run(cfg: Config) -> io::Result<()> {
    let cfg = cfg.sanitized();
    let frames = frame_set(cfg.emoji);

    let running = Arc::new(AtomicBool::new(true));
    {
        let running = running.clone();
        ctrlc::set_handler(move || {
            running.store(false, Ordering::SeqCst);
        })
        .ok();
    }

    print!("\x1b[?25l");
    println!("风扇旋转中（Ctrl+C 退出）...");

    let mut idx = 0usize;
    let delay = Duration::from_millis(cfg.speed_ms);

    while running.load(Ordering::SeqCst) {
        std::thread::sleep(delay);
        print!("\r\x1b[2K{}", frames[idx % frames.len()]);
        io::stdout().flush().ok();
        idx = idx.wrapping_add(1);
    }

    print!("\r\x1b[2K\x1b[?25h\n");
    Ok(())
}

pub fn frame_set(emoji: bool) -> &'static [&'static str] {
    if emoji {
        &EMOJI_FRAMES
    } else {
        &ASCII_FRAMES
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn frame_selection() {
        assert_eq!(frame_set(false)[0], "-");
        assert_eq!(frame_set(true)[0], "◐");
    }
}
