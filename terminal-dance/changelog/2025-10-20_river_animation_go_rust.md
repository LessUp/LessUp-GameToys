# 2025-10-20

- 新增 `CodingToys/river/go/`：`go.mod`、`main.go`，实现终端小溪动画（支持 `-emoji`）。
- 新增 `CodingToys/river/rust/`：`Cargo.toml`、`src/main.rs`，实现终端小溪动画（支持 `-emoji`）。
- 设计遵循 KISS；不引入多余依赖（Rust 仅使用 `ctrlc` 用于优雅退出）。
- 修复 Go 帧率控制：改为阻塞等待 ticker，`-speed` 现在严格生效。
- 调整为横向流动（Go 与 Rust）：河道按列变化、水平滚动；emoji 模式表现优化。
