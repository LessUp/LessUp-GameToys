# CodingToys

## 项目简介
- **目标**：收集一系列轻量的交互式小玩具与动画示例，优先原生技术与可直接运行的体验。
- **原则**：遵循 KISS，保持单个模块独立、易读、易复用。

## 目录速览
- **`webfront/`**：前端可视化与动画示例，每个子目录通常包含单文件 `index.html` 及需求说明。
- **`terminal-dance/`**：终端动画合集，提供 Go 与 Rust 两种实现，可通过命令行参数控制表现。
- **`changelog/`**：逐日记录各模块的新增、修改与修正，便于追踪历史。
- **`tools/`**：开发辅助脚本集合，包含从元数据生成首页等自动化任务。
- **`docs/`**：项目规范、设计语言与协作流程文档。

## 使用方式
- **Web 模块**：进入对应子目录，直接在浏览器打开 `index.html` 即可体验。
- **终端模块**：阅读子目录 README 或源码注释，使用 Go/Rust 运行对应程序，按需传入参数。
- **变更追踪**：参考 `changelog/` 下的 Markdown 文件了解最近改动。

## 约定
- **语言**：文档与提示默认使用中文。
- **规范**：提交前请先阅读《[项目规范与设计语言](docs/PROJECT_STANDARDS.md)》与《[贡献指南](CONTRIBUTING.md)》。
- **记录**：每次修改需同步更新 `changelog/`，保持历史清晰。
- **依赖**：除非必要，不引入额外依赖，确保示例可以即取即用。

## 模块元数据与自动化
- 每个 Web/终端模块目录下需维护一份 `meta.json`，描述标题、简介、分类、图标等基础信息；终端模块额外列出常用命令参数。
- `tags` 字段可选择性指定样式变体（如 `physics`、`animation`、`interactive`、`game`），用于统一的标签配色。
- 示例结构：

```json
{
  "title": "模块标题",
  "description": "一句话描述",
  "category": "所属分类",
  "icon": "🔥",
  "tags": [
    { "label": "物理引擎", "variant": "physics" }
  ],
  "order": 10
}
```

- 首页卡片通过 `python tools/sync_webfront_index.py` 自动生成；新增或调整模块后，请先更新对应 `meta.json`，再运行脚本同步 `webfront/index.html`。
