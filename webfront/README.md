# webfront 前端可视化示例

单文件 HTML/CSS/JS 实现的交互式前端小项目，无需构建工具，直接在浏览器打开即可体验。

## 模块清单

### 物理模拟
- `Free-Fall-of-Small-Balls-in-a-Hexagon/`：六边形容器内的小球重力与碰撞演示。
- `Double-Pendulum/`：可调参数的双摆混沌系统。
- `Particle-System/`：多力场粒子系统引擎。
- `Wave-Simulation/`：高度场水波传播模拟。
- `Spring-Physics/`：弹簧-质点布料动力学。

### 视觉特效
- `Particle-Effects/`：烟花、能量脉冲等粒子爆炸特效。
- `Morphing-Shapes/`：几何形状的平滑变换动画。
- `3D-Transforms/`：CSS 与 Canvas 的 3D 变换演示。
- `weather/`：动态天气卡片集合。

### 3D 交互
- `Cube-Simulation/`：3×3 魔方交互模拟器。

### 互动小游戏
- `2048/`：经典 2048 合并玩法。
- `snake/`：霓虹风格贪吃蛇。

## 元数据规范
- 每个模块目录包含 `meta.json`，字段覆盖标题、描述、分类、图标、标签与排序权重。
- `tags.variant` 支持 `physics`、`animation`、`interactive`、`game` 等样式，便于统一配色。
- 元数据与模块 README 同步更新，可减少首页维护成本。

## 首页同步流程
1. 更新或新增模块时，先维护对应 `meta.json` 与 README。
2. 运行 `python ../tools/sync_webfront_index.py`（或在仓库根目录执行 `python tools/sync_webfront_index.py`）。
3. 提交自动生成的 `webfront/index.html` 以确保首页卡片与目录保持一致。

## 使用方式
进入对应子目录，直接在浏览器中打开 `index.html` 即可体验。

## 设计原则
- 保持单文件结构，无外部依赖。
- 使用原生 HTML/CSS/JavaScript。
- 追求简洁、可读、可复用的代码风格。
- 支持桌面与移动端响应式设计。
