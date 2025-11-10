# terminal-dance

终端动画与轻量小游戏合集（KISS）：河流、火车、旋转风扇（fun）、跑动小猫（cat）、贪吃蛇（snake）、康威生命游戏（life）。支持 ASCII 与 Emoji 两种风格（按需）。

## 模块清单
- `river/`：波动的河道动画。
- `train/`：持续行进的小火车，支持调节车厢与帧率。
- `fun/`：旋转风扇动画，演示最小帧动画结构。
- `cat/`：奔跑的小猫，展示横向滚动与宽度控制。
- `snake/`：贪吃蛇小游戏，支持键盘操控。
- `life/`：康威生命游戏，支持随机密度与 Emoji 渲染。

每个目录均包含 `meta.json`，用于记录模块介绍、图标以及支持的命令行参数。

## 统一参数规范
| 参数 | Go 标记 | Rust 标记 | 默认值 | 适用模块 | 说明 |
| --- | --- | --- | --- | --- | --- |
| `emoji` | `-emoji` | `--emoji` | `false` | 全部 | 是否使用 Emoji 渲染帧。 |
| `speed` | `-speed` | `--speed` | 模块各异 | 全部 | 每帧延迟（毫秒）或帧率，数值越小越快。 |
| `w` | `-w` | `--w` | 模块各异 | `cat`、`snake`、`life`、`river` | 渲染宽度或世界宽度。缺省读取终端列数。 |
| `h` | `-h` | `--h` | 模块各异 | `snake`、`life`、`river` | 渲染高度或世界高度。 |
| `density` | `-density` | `--density` | `0.3` | `life` | 初始细胞密度（0.0-1.0）。 |
| `rw` | `-rw` | `-rw` | `8` | `river` | 河道宽度（字符行数）。Emoji 模式会自动收窄。 |
| `cars` | `-cars` | `--cars` | `3` | `train` | 车厢数量（0-20）。 |

参数默认值与更多说明请参阅各模块的 `meta.json`。若新增参数，请同步更新表格保持一致。

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
- 河流：
```bash
cd river/go
go run ./ -emoji=false -speed=60 -w=80 -h=24 -rw=8
```
- 火车：
```bash
cd train/go
go run ./ -emoji=false -speed=12 -cars=3
```

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
- 河流：
```bash
cd river/rust
cargo run --release -- -emoji=false -speed=60 -w=80 -h=24 -rw=8
```
- 火车：
```bash
cd train/rust
cargo run --release -- --emoji=false --speed=12 --cars=3
```

## 交互
- 各动画/游戏均支持 `Ctrl+C` 退出，会恢复光标显示。
- 贪吃蛇：使用 `WASD` 或方向键控制方向，`Q` 键退出。
- 生命游戏：自动演化，`Ctrl+C` 退出。

## 备注
- 终端与字体对 Emoji 宽度的处理不完全一致，排版可能略有差异。
- 保持 KISS 设计，尽量少依赖，跨平台 ANSI 控制序列（在 Linux 下表现良好）。
