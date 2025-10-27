# terminal-dance

终端动画与轻量小游戏合集（KISS）：河流、火车、风扇（fun）、小猫（cat）、贪吃蛇（snake）、康威生命游戏（life）。支持 ASCII 与 Emoji 两种风格（按需）。

## 目录结构

- `river/` 河流动画（Go/Rust）
- `train/` 小火车动画（Go/Rust）
- `fun/` 旋转风扇（Go/Rust）
- `cat/` 跑动小猫（Go/Rust）
- `snake/` 贪吃蛇游戏（Go/Rust）
- `life/` 康威生命游戏（Go/Rust）

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
- 贪吃蛇：
```bash
cd snake/go
go run ./ -emoji=false -speed=150 -w=20 -h=10
```
- 生命游戏：
```bash
cd life/go
go run ./ -emoji=false -speed=100 -w=40 -h=20 -density=0.3
```

参数说明（Go）：
- `-emoji` 是否使用 Emoji 风格
- `-speed` 每帧延迟(毫秒)，数值越小越快
- `-w` 宽度参数（小猫/贪吃蛇/生命游戏）
- `-h` 高度参数（贪吃蛇/生命游戏）
- `-density` 生命游戏初始细胞密度(0.0-1.0)

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
- 贪吃蛇：
```bash
cd snake/rust
cargo run --release -- --emoji=false --speed=150 --w=20 --h=10
```
- 生命游戏：
```bash
cd life/rust
cargo run --release -- --emoji=false --speed=100 --w=40 --h=20 --density=0.3
```

参数说明（Rust）：
- `--emoji` 是否使用 Emoji 风格
- `--speed` 每帧延迟(毫秒)，数值越小越快
- `--w` 宽度参数（小猫/贪吃蛇/生命游戏）
- `--h` 高度参数（贪吃蛇/生命游戏）
- `--density` 生命游戏初始细胞密度(0.0-1.0)

## 交互
- 各动画均支持 `Ctrl+C` 退出，会恢复光标显示。
- 贪吃蛇：使用 `WASD` 或方向键控制方向，`Q` 键退出。
- 生命游戏：自动演化，`Ctrl+C` 退出。

## 备注
- 终端与字体对 Emoji 宽度的处理不完全一致，排版可能略有差异。
- 保持 KISS 设计，尽量少依赖，跨平台 ANSI 控制序列（在 Linux 下表现良好）。
