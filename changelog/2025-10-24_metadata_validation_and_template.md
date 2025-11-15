# 2025-10-24 元数据校验与模板更新

## 变更摘要
- 为前端与终端模块补齐元数据字段，增加更新日期与关联变更记录。
- 扩展同步脚本支持校验模式与变更日志关联检查，保障 CI 可覆盖。
- 更新变更日志模板，强调摘要、影响范围与 meta 校验确认。

## 影响范围
- `webfront/*/meta.json`
- `terminal-dance/*/meta.json`
- `tools/sync_webfront_index.py`
- `changelog/_template.md`

## Meta 校验
- [x] 已更新关联模块的 `meta.json`
- [x] 已运行 `python tools/sync_webfront_index.py --check`

## 详情
### 新增
- 新增 `2025-10-24_metadata_validation_and_template.md`，用于记录本次元数据校验策略。

### 修改
- 扩展同步脚本，新增 `--check` 模式并执行 changelog 关联校验。
- 补充所有模块的 `updated_at` 与 `changelog` 字段，并设定 `terminal-dance` 模块排序。
- 调整 `changelog/_template.md` 结构，要求记录摘要、影响范围与 meta 校验步骤。

### 修复
- 无。

### 备注
- CI 应执行 `python tools/sync_webfront_index.py --check` 以验证索引和变更日志一致性。
