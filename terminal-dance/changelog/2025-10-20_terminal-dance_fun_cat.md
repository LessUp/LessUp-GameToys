# 2025-10-20 terminal-dance: 新增风扇与小猫

- 新增目录：`fun/go/`, `fun/rust/`, `cat/go/`, `cat/rust/`
- 新增 Go：
  - `fun/go/main.go`：旋转风扇（ASCII/Emoji），参数 `--emoji --speed`
  - `cat/go/main.go`：跑动小猫（ASCII/Emoji），参数 `--emoji --speed --w`
- 新增 Rust：
  - `fun/rust/`：`Cargo.toml`, `src/main.rs`（旋转风扇，支持 Ctrl+C 恢复光标）
  - `cat/rust/`：`Cargo.toml`, `src/main.rs`（跑动小猫，支持 Ctrl+C 恢复光标）
- 新增 `README.md`：统一说明构建与运行、参数
- 补充 `.gitignore`（忽略子 crate 的 `**/target/` 及常见构建产物）
