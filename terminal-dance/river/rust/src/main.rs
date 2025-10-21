use std::env;
use std::io::{self, Write};
use std::sync::atomic::{AtomicBool, Ordering};
use std::sync::Arc;
use std::{thread, time};

fn main() {
    let mut w: usize = 80;
    let mut h: usize = 24;
    let mut emoji = false;
    let mut speed: u64 = 60;
    let mut rw: usize = 8;

    let mut args = env::args().skip(1);
    while let Some(arg) = args.next() {
        match arg.as_str() {
            "-w" => if let Some(v) = args.next() { w = v.parse().unwrap_or(w); },
            "-h" => if let Some(v) = args.next() { h = v.parse().unwrap_or(h); },
            "-emoji" => { emoji = true; },
            "-speed" => if let Some(v) = args.next() { speed = v.parse().unwrap_or(speed); },
            "-rw" => if let Some(v) = args.next() { rw = v.parse().unwrap_or(rw); },
            _ => {}
        }
    }

    if emoji && rw > 4 { rw = 4; }
    if rw >= h { rw = h.saturating_sub(1); }

    let running = Arc::new(AtomicBool::new(true));
    let r = running.clone();
    let _ = ctrlc::set_handler(move || { r.store(false, Ordering::SeqCst); });

    print!("\x1b[?25l");
    let _ = io::stdout().flush();

    let freq: f64 = 0.20;
    let mut frame: usize = 0;

    while running.load(Ordering::SeqCst) {
        let mut screen = String::new();
        screen.push_str("\x1b[H\x1b[2J");
        // 横向河道：按列计算中心行
        let half_h: isize = (rw / 2) as isize;
        let mut amp: isize = (((h.saturating_sub(rw).saturating_sub(1)) as f64) * 0.25) as isize;
        if amp < 1 { amp = 1; }
        for row in 0..h {
            for col in 0..w {
                let f = (col as f64) * freq;
                let s = f.sin();
                let center: isize = (h as isize) / 2 + ((amp as f64) * s) as isize;
                let r: isize = row as isize;
                let in_river = r >= (center - half_h) && r < (center + half_h);
                if in_river {
                    if emoji {
                        let idx = (((col as isize - frame as isize) % 3) + 3) % 3;
                        let ch = match idx { 0 => "🌊", 1 => "💦", _ => "💧" };
                        screen.push_str(ch);
                    } else {
                        let idx = (((col as isize - frame as isize) % 3) + 3) % 3;
                        let ch = match idx { 0 => '~', 1 => '-', _ => '=' };
                        screen.push(ch);
                    }
                } else {
                    screen.push(' ');
                }
            }
            screen.push('\n');
        }
        print!("{}", screen);
        let _ = io::stdout().flush();
        frame = frame.wrapping_add(1);
        thread::sleep(time::Duration::from_millis(speed));
    }

    print!("\x1b[?25h\x1b[0m\n");
    let _ = io::stdout().flush();
}
