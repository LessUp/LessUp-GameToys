# 2025-10-21 Webfront 3D Rubik's Cube (3x3)

- 路径：`CodingToys/webfront/3D-Rubik's-Cube-Simulation/index.html`
- 设计：单文件、零依赖（原生 HTML/CSS/JS）、KISS。
- 渲染：CSS 3D Transforms + 透视；分层动画采用临时包装容器过渡，结束后更新状态并复位 DOM。
- 功能：
  - 鼠标拖拽旋转视角、滚轮缩放、双击复位。
  - 按钮操作 U/D/L/R/F/B 及其逆时针（'）。
  - 打乱（25 步随机）与重置。
- 实现要点：
  - 数据模型为 26 个 Cubie（省略中心空位），每个 Cubie 存储 `pos`（-1/0/1）与 `orient`（本地轴→世界轴）以及本地贴纸集合。
  - 旋转时：选取对应层，使用包装容器进行 CSS 旋转动画；动画结束后应用逻辑变换 `rotatePos`/`rotateWorldVec` 并重新计算可见面与贴纸颜色。
  - 贴纸颜色绑定 Cubie 本地面，随 `orient` 旋转，不做全局纹理映射。
- 使用：直接用浏览器打开 `index.html` 即可。

---

- 修正（2025-10-21）：
  - 修复 `rotateWorldVec()` 绕 Y/Z 轴的符号映射，确保与右手系一致：
    - `R_y(+90): x -> z, z -> -x`
    - `R_z(+90): x -> -y, y -> x`
