# terminal-dance

终端小动画合集（KISS）：河流、火车、风扇（fun）、小猫（cat）。支持 ASCII 与 Emoji 两种风格（按需）。

## 目录结构

- `river/` 河流动画（Go/Rust，已存在）
- `train/` 小火车动画（Go/Rust，已存在）
- `fun/` 旋转风扇（Go/Rust，本次新增）
- `cat/` 跑动小猫（Go/Rust，本次新增）

## 构建与运行

### Go 版本

- 风扇：
```bash
cd fun/go
go run ./ -emoji=false -speed=80
```
- 小猫：
```bash
cd cat/go
go run ./ -emoji=false -speed=80 -w=80
```

参数说明（Go）：
- `-emoji` 是否使用 Emoji 风格
- `-speed` 每帧延迟(毫秒)，数值越小越快
- `-w` 小猫横向可用宽度（默认读取 `COLUMNS`，无则 80）

### Rust 版本

- 风扇：
```bash
cd fun/rust
cargo run --release -- --emoji=false --speed=80
```
- 小猫：
```bash
cd cat/rust
cargo run --release -- --emoji=false --speed=80 --w=80
```

参数说明（Rust）：
- `--emoji` 是否使用 Emoji 风格
- `--speed` 每帧延迟(毫秒)，数值越小越快
- `--w` 小猫横向可用宽度

## 交互
- 各动画均支持 `Ctrl+C` 退出，会恢复光标显示。

## 备注
- 终端与字体对 Emoji 宽度的处理不完全一致，排版可能略有差异。
- 保持 KISS 设计，尽量少依赖，跨平台 ANSI 控制序列（在 Linux 下表现良好）。
