# 贡献指南

## 本地验证流程

在提交变更之前，请确保本地通过以下检查：

1. 安装依赖：
   - Node.js：在仓库根目录执行 `npm install`，安装 `package.json` 中声明的前端工具。
   - Python：使用 `pip install black ruff` 安装格式化与 Lint 工具。
   - Go：确保安装官方工具链，并使用 `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` 获取 `golangci-lint`。
   - Rust：通过 `rustup component add rustfmt clippy` 安装 `cargo fmt` 与 `cargo clippy` 所需组件。
2. 运行统一检查：
   - 执行 `make ci` 依次调用格式化、Lint 与核心测试命令。

### 关键命令说明

| 命令 | 作用 | 常见修复命令 |
| --- | --- | --- |
| `npm run prettier` | 使用 Prettier 校验 `webfront/` 与 `tools/` 下的前端相关资源格式。 | `npm run prettier:fix` 自动修复可处理的问题。 |
| `npm run eslint` | 使用 ESLint 对 `webfront/` 及相关脚本进行静态检查。 | 根据提示调整代码或配置。 |
| `npm run black` | 使用 Black 校验 `tools/` 下 Python 代码的格式。 | `npm run black:fix` 自动格式化。 |
| `npm run ruff` | 使用 Ruff 对 `tools/` 下 Python 代码进行快速 Lint。 | `npm run ruff:fix` 自动修复 Ruff 支持的问题。 |
| `make gofmt` | 确认 `terminal-dance/` 内所有 Go 代码经过 `gofmt`。 | 在各 Go 子项目中运行 `gofmt -w <file>`。 |
| `make golangci` | 在每个 Go 子项目中执行 `golangci-lint run`。 | 修复提示或更新 `golangci-lint` 配置。 |
| `make go-test` | 对 Go 子项目执行 `go test ./...`。 | 编写/修复测试并保证依赖可用。 |
| `make cargo-fmt` | 在每个 Rust 子项目中执行 `cargo fmt -- --check`。 | 使用 `cargo fmt` 自动格式化。 |
| `make cargo-clippy` | 在每个 Rust 子项目中执行 `cargo clippy --all-targets --all-features -- -D warnings`。 | 根据提示修复代码或适当添加 `#[allow]`。 |
| `make cargo-test` | 在每个 Rust 子项目中执行 `cargo test`。 | 修复失败的测试或缺失依赖。 |

### 失败排查建议

- 如遇格式化或 Lint 失败，优先运行上表中的修复命令，并根据输出定位具体文件。
- 若依赖缺失导致命令无法执行，请按照“安装依赖”步骤重新安装，必要时检查版本是否满足要求。
- Go 或 Rust 项目可能需要额外的系统依赖（例如 `pkg-config`、本地 C 库）。请根据命令输出提示补充。
- 如果 `make ci` 某一步骤失败，可单独运行对应命令以快速定位问题，确认修复后再执行一次 `make ci` 确认全部通过。
