# Contributing Guide

感谢你对 LessUp GameToys 的关注！为了保持项目高质量、可维护与一致的体验，请在贡献前阅读下列指南。

## 快速上手
1. Fork 仓库并创建特性分支：`feature/<topic>` 或 `fix/<topic>`。
2. 安装必要依赖：Go、Rust、Node.js、Python 3.9+。
3. 在提交代码前执行对应语言的格式化与静态检查（见 [项目规范](docs/PROJECT_STANDARDS.md)）。

## 提交流程
- 使用 [Conventional Commits](https://www.conventionalcommits.org/zh-hans/v1.0.0/) 书写提交信息，如 `feat: add hexagon simulation controls`。
- 每个 PR 需包含：
  - 变更摘要与动机。
  - 测试或演示结果（命令输出、截图或录屏）。
  - 关联 issue（如有）。
- 确保更新 `changelog/` 中对应日期的记录或新增条目。

## 代码规范速查
- **前端**：遵循 `docs/PROJECT_STANDARDS.md` 中的设计语言，使用 ES Modules、语义化 HTML 与 CSS Design Tokens。
- **终端动画**：Go/Rust 主函数保持精简，将业务逻辑拆分到独立包/模块，提供单元测试。
- **脚本/工具**：提供 `--help` 与 `--dry-run`，输出结构化日志。
- **文档**：模块必须具备 `README.md` 与 `meta.json`，保持描述同步。

## 评审标准
- 代码通过格式化、lint 与测试。
- 动画表现平稳（目标 60 FPS）且具备性能说明。
- 文档清晰描述使用方法、参数、已知限制。
- 设计符合规范，新增组件经过评审流程。

遵循以上流程有助于我们快速评审并合并你的贡献，感谢支持！
