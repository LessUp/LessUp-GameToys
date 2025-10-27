# Snake Neon 霓虹贪吃蛇

Canvas 实现的霓虹风格贪吃蛇小游戏，支持键盘与触摸双操作模式。

## 特性
- Canvas 动态渲染，霓虹渐变效果
- 键盘方向键 / WASD 控制
- 移动端虚拟按键 + 触摸滑动
- 自带音效反馈（使用 Web Audio API）
- 本地最高分记录（LocalStorage）
- 自适应速度调节，越吃越快

## 使用方式
在浏览器中打开 `index.html` 即可体验。

## 操作说明
- **桌面端**：方向键或 WASD 控制移动
- **移动端**：滑动屏幕或点击虚拟按键
- 吃到能量球可加分并延长身体，撞墙或撞到自己会重新开始

## 技术要点
- 使用 `requestAnimationFrame` 驱动游戏循环
- Web Audio API 生成简单音效
- LocalStorage 存储最高分
- 自定义虚拟按键与触摸手势兼容
