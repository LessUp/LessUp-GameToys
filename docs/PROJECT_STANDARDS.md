# 项目规范与设计语言

本规范旨在为 LessUp GameToys 项目的前端、终端动画、脚本与文档提供统一的设计语言与协作流程。所有新模块与修改需遵循以下原则，以确保一致的体验与高质量交付。

## 1. 架构设计原则

- **模块化与独立部署**：每个示例/玩具保持最小依赖，可独立运行与演示。
- **可复用组件**：公共逻辑、样式与工具函数应抽离到共享模块或 `tools/`，避免重复实现。
- **渐进增强**：默认使用原生技术，必要时才引入第三方库，并在文档中标注理由与版本。
- **可观测性**：关键交互、动画参数与性能瓶颈应通过注释或调试面板暴露，便于调优。

## 2. 前端规范

### 2.1 结构与命名
- 目录统一采用 `kebab-case`，HTML/CSS/JS 文件命名对齐模块标题。
- 同一模块中的资源按类型分层：`index.html`、`styles/`、`scripts/`、`assets/`。
- 自定义元素或组件使用 `gt-` 前缀（GameToys 缩写）区分。

### 2.2 HTML
- 使用语义化标签，确保键盘可访问性。
- 所有交互元素需具备 `aria-*` 属性，确保屏幕阅读器支持。
- 元数据 (`meta.json`) 与页面信息保持同步，描述应在 80 字内。

### 2.3 CSS/设计语言
- 采用 **Design Tokens** 统一视觉：
  - 颜色：基础色（`--color-surface`、`--color-text`、`--color-accent`）。
  - 间距：`--space-xs/s/m/l/xl`，遵循 1.5 倍数节奏。
  - 字体：优先 `"JetBrains Mono", "Fira Code", monospace`。
- 使用 CSS 自定义属性集中管理主题，禁止硬编码颜色/尺寸。
- 动画默认曲线：`cubic-bezier(0.4, 0, 0.2, 1)`，时长 `250ms` 起步。

### 2.4 JavaScript
- 遵循 ES Modules，所有脚本以 `type="module"` 引入。
- 使用命名导出，函数名动词开头，如 `initScene`、`renderFrame`。
- 与 DOM 交互需封装在 `setup*` 或 `bind*` 函数中，便于测试。
- 提供 `destroy` 钩子以释放事件监听、动画帧。

### 2.5 可测试性
- 在 `scripts/` 下提供 `__tests__/` 或文档中说明验证步骤。
- 渲染/模拟逻辑应与 UI 分离，便于在 Node 环境中做单元测试。

## 3. 终端动画（Go/Rust）规范

- **入口结构**：主函数仅负责参数解析与调用 `Run()` 或 `Start()`。
- **配置加载**：将可调参数集中在结构体中，并提供默认值与命令行覆盖。
- **错误处理**：Go 使用 `fmt.Errorf("context: %w", err)` 包装；Rust 使用 `anyhow::Result` 或 `thiserror`。
- **渲染循环**：抽象为 `update` 与 `draw` 两阶段，确保可移植性。
- **测试**：关键算法放入独立包/模块，覆盖单元测试或基准测试。

## 4. 脚本与自动化

- Python/Node 脚本需提供 `--dry-run` 选项，默认读取配置文件。
- 代码须通过 `black`/`ruff` 或 `eslint`/`prettier` 等工具格式化（见下表）。
- 输出需具备详细日志等级（`info/warn/error`）。

| 语言 | 格式化 | 静态检查 |
| ---- | ------ | -------- |
| JavaScript/TypeScript | `prettier --write` | `eslint .` |
| Go | `gofmt` | `golangci-lint run` |
| Rust | `cargo fmt` | `cargo clippy --all-targets --all-features` |
| Python | `black` | `ruff check` |

## 5. 文档与元数据

- 每个模块必须包含 `README.md`，包含：目标、运行方式、关键参数、截图（如适用）。
- `meta.json` 字段需通过 `tools/sync_webfront_index.py` 验证，保证首页一致性。
- 变更需在 `changelog/` 目录新增条目，格式遵循模板 `_template.md`。

## 6. 版本控制与提交规范

- Commit 信息遵循 [Conventional Commits](https://www.conventionalcommits.org/zh-hans/v1.0.0/)。
- PR 描述需概括动机、变更点、验证方式，并关联相关 issue。
- 禁止在主分支进行 force push；功能分支命名 `feature/<topic>`，修复分支 `fix/<topic>`。

## 7. 质量基线

- 所有代码需通过 lint、格式化与单元测试方可合并。
- 对于动画性能，目标保持 60 FPS；若达不到需在文档中解释并提供优化计划。
- 安全审查：前端避免使用 `innerHTML` 注入；脚本加载外部资源需做校验与异常处理。

## 8. 设计评审流程

1. 在 Issue 中描述新模块/改动的目标、交互流程与技术栈。
2. 附上线框、动效示意或伪代码，供团队讨论。
3. 通过评审后开始开发，过程中保持 Changelog 更新。
4. 提交 PR 时附带演示视频/GIF 或命令截图，确认验收标准。

---

维护者可根据实际需求迭代本规范；重大调整需在 `changelog/` 中同步公告。
