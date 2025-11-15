use clap::Parser;

#[derive(Parser, Debug, Clone)]
#[command(author, version, about = "旋转风扇 (ASCII/Emoji)")]
pub struct Config {
    #[arg(long, default_value_t = false, help = "使用表情风格")]
    pub emoji: bool,
    #[arg(long, default_value_t = 80u64, help = "每帧延迟(毫秒)")]
    pub speed_ms: u64,
}

impl Config {
    pub fn parse() -> Result<Self, clap::Error> {
        Self::try_parse()
    }

    pub fn sanitized(mut self) -> Self {
        if self.speed_ms == 0 {
            self.speed_ms = 1;
        }
        self
    }
}
