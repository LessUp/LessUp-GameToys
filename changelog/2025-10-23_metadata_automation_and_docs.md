# 2025-10-23 元数据规范与文档自动化

## 新增
- 为 `webfront/` 与 `terminal-dance/` 下的每个模块补充 `meta.json`，统一记录标题、简介、图标、分类以及参数信息。
- 创建 `tools/sync_webfront_index.py`，可根据元数据生成首页卡片，新增 `.tag.game` 样式支持互动类标签。
- 在 `webfront/` 各演示目录补充 README，覆盖双摆、粒子系统、布料、波纹、变形动画、魔方、天气等模块的使用说明与交互提示。
- 新增 `changelog/_template.md`，用于约束后续变更记录结构。

## 修改
- `webfront/index.html` 由脚本自动生成，新增互动小游戏分类并引用新的标签配色。
- 更新根目录 README，说明 `meta.json` 规范及首页同步流程；完善 `webfront/README.md` 与 `terminal-dance/README.md`，补充模块清单、元数据说明及统一参数表。
- 2048 与贪吃蛇标签改用 `game` 变体，便于与互动类演示区分。

## 备注
- 更新或新增模块时，请优先维护对应 `meta.json` 与 README，再执行 `python tools/sync_webfront_index.py` 以保持首页一致。
