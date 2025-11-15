use clap::Parser;

#[derive(Parser, Debug, Clone)]
#[command(author, version, about = "康威生命游戏 (ASCII/Emoji)")]
pub struct Config {
    #[arg(long, default_value_t = 40usize, help = "世界宽度")]
    pub width: usize,
    #[arg(long, default_value_t = 20usize, help = "世界高度")]
    pub height: usize,
    #[arg(long, default_value_t = false, help = "使用表情风格")]
    pub emoji: bool,
    #[arg(long, default_value_t = 100u64, help = "每帧延迟(毫秒)")]
    pub speed_ms: u64,
    #[arg(long, default_value_t = 0.3f64, help = "初始密度(0.0-1.0)")]
    pub density: f64,
}

impl Config {
    pub fn parse() -> Result<Self, clap::Error> {
        Self::try_parse()
    }

    pub fn sanitized(mut self) -> Self {
        self.width = self.width.max(4);
        self.height = self.height.max(4);
        self.density = self.density.clamp(0.0, 1.0);
        self.speed_ms = self.speed_ms.max(30);
        self
    }
}
