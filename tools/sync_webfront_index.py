#!/usr/bin/env python3
"""Generate webfront/index.html from module metadata."""
from __future__ import annotations

import json
from collections import defaultdict
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
WEBFRONT_DIR = ROOT / "webfront"
INDEX_PATH = WEBFRONT_DIR / "index.html"

CATEGORY_ORDER = {
    "物理模拟": 10,
    "视觉特效": 20,
    "3D 交互": 30,
    "互动小游戏": 40,
}

def load_metadata() -> list[dict]:
    modules = []
    for meta_path in sorted(WEBFRONT_DIR.glob("*/meta.json")):
        with meta_path.open("r", encoding="utf-8") as fh:
            data = json.load(fh)
        required = ["title", "description", "category", "icon"]
        missing = [key for key in required if not data.get(key)]
        if missing:
            raise SystemExit(f"Missing keys {missing} in {meta_path}")
        data.setdefault("tags", [])
        data.setdefault("order", 0)
        data["slug"] = meta_path.parent.name
        modules.append(data)
    return modules

def build_cards(modules: list[dict]) -> str:
    grouped: dict[str, list[dict]] = defaultdict(list)
    for module in modules:
        grouped[module["category"]].append(module)


    cards_output: list[str] = []
    for category in sorted(grouped.keys(), key=lambda c: (CATEGORY_ORDER.get(c, 999), c)):
        cards_output.append(f"    <h2 class=\"category-title\">{category}</h2>")
        cards_output.append("    <div class=\"grid\">")
        for module in sorted(grouped[category], key=lambda m: (int(m.get("order", 0)), m["title"])):
            tags_lines = []
            for tag in module.get("tags", []):
                label = tag["label"] if isinstance(tag, dict) else str(tag)
                if isinstance(tag, dict):
                    variant = tag.get("variant", "").strip()
                else:
                    variant = ""
                class_name = "tag"
                if variant:
                    class_name += f" {variant}"
                tags_lines.append(f"          <span class=\"{class_name}\">{label}</span>")
            if tags_lines:
                tags_block = ["          <div class=\"card-tags\">"] + tags_lines + ["          </div>"]
            else:
                tags_block = ["          <div class=\"card-tags\"></div>"]
            card_lines = [
                f"      <a href=\"{module['slug']}/\" class=\"card\">",
                "        <div class=\"card-preview\">",
                f"          <div class=\"card-icon\">{module['icon']}</div>",
                "        </div>",
                "        <div class=\"card-content\">",
                f"          <h3 class=\"card-title\">{module['title']}</h3>",
                f"          <p class=\"card-description\">{module['description']}</p>",
                *tags_block,
                "        </div>",
                "      </a>",
            ]
            cards_output.append("\n".join(card_lines))
            cards_output.append("")
        cards_output.append("    </div>")
        cards_output.append("")
    return "\n".join(cards_output).rstrip()

HTML_TEMPLATE = """<!DOCTYPE html>
<html lang=\"zh-CN\">
<head>
  <meta charset=\"utf-8\" />
  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" />
  <title>Webfront · 交互式前端演示集</title>
  <style>
    :root {
      --bg-dark: #0a0e14;
      --bg-mid: #0f1419;
      --bg-light: #1a1f29;
      --text: #e6eef8;
      --text-muted: #8b95a8;
      --accent-1: #4ea1ff;
      --accent-2: #a855f7;
      --accent-3: #ec4899;
      --accent-4: #f59e0b;
      --accent-5: #38bdf8;
      --card-bg: rgba(26, 31, 41, 0.6);
      --card-border: rgba(78, 161, 255, 0.15);
      --shadow: 0 20px 60px rgba(0, 0, 0, 0.4);
    }

    * { box-sizing: border-box; margin: 0; padding: 0; }

    html, body {
      height: 100%;
      background: var(--bg-dark);
      color: var(--text);
      font-family: system-ui, -apple-system, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, 'PingFang SC', 'Noto Sans CJK SC', 'Microsoft YaHei', sans-serif;
      overflow-x: hidden;
    }

    body {
      background:
        radial-gradient(circle at 20% 10%, rgba(78, 161, 255, 0.08) 0%, transparent 40%),
        radial-gradient(circle at 80% 80%, rgba(168, 85, 247, 0.08) 0%, transparent 40%),
        linear-gradient(135deg, var(--bg-dark) 0%, var(--bg-mid) 50%, var(--bg-dark) 100%);
      position: relative;
    }

    /* Animated background particles */
    .bg-canvas {
      position: fixed;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      z-index: 0;
      opacity: 0.4;
    }

    .container {
      position: relative;
      z-index: 1;
      max-width: 1400px;
      margin: 0 auto;
      padding: 40px 24px;
    }

    /* Header */
    header {
      text-align: center;
      padding: 60px 0 80px;
      position: relative;
    }

    h1 {
      font-size: clamp(2.5rem, 5vw, 4rem);
      font-weight: 800;
      background: linear-gradient(135deg, var(--accent-1) 0%, var(--accent-2) 50%, var(--accent-3) 100%);
      -webkit-background-clip: text;
      background-clip: text;
      -webkit-text-fill-color: transparent;
      margin-bottom: 16px;
      letter-spacing: -0.02em;
      animation: fadeInUp 0.8s ease-out;
    }

    .subtitle {
      font-size: clamp(1rem, 2vw, 1.25rem);
      color: var(--text-muted);
      max-width: 600px;
      margin: 0 auto 32px;
      line-height: 1.6;
      animation: fadeInUp 0.8s ease-out 0.1s both;
    }

    .badge {
      display: inline-flex;
      align-items: center;
      gap: 8px;
      padding: 8px 20px;
      background: var(--card-bg);
      border: 1px solid var(--card-border);
      border-radius: 999px;
      font-size: 0.875rem;
      color: var(--accent-1);
      backdrop-filter: blur(10px);
      animation: fadeInUp 0.8s ease-out 0.2s both;
    }

    .badge::before {
      content: '';
      width: 8px;
      height: 8px;
      background: var(--accent-1);
      border-radius: 50%;
      animation: pulse 2s ease-in-out infinite;
    }

    /* Grid */
    .grid {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(340px, 1fr));
      gap: 28px;
      margin-bottom: 60px;
    }

    @media (max-width: 768px) {
      .grid {
        grid-template-columns: 1fr;
      }
    }

    /* Category Title */
    .category-title {
      font-size: 1.75rem;
      font-weight: 700;
      color: var(--text);
      margin: 60px 0 24px;
      padding-bottom: 12px;
      border-bottom: 2px solid var(--card-border);
      display: flex;
      align-items: center;
      gap: 12px;
    }

    .category-title::before {
      content: '';
      width: 4px;
      height: 28px;
      background: linear-gradient(180deg, var(--accent-1), var(--accent-2));
      border-radius: 2px;
    }

    /* Card */
    .card {
      position: relative;
      background: var(--card-bg);
      border: 1px solid var(--card-border);
      border-radius: 20px;
      overflow: hidden;
      transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
      backdrop-filter: blur(10px);
      animation: fadeInUp 0.6s ease-out both;
      cursor: pointer;
      text-decoration: none;
      color: inherit;
      display: block;
    }

    .card::before {
      content: '';
      position: absolute;
      inset: 0;
      background: linear-gradient(135deg, rgba(78, 161, 255, 0.15), rgba(236, 72, 153, 0.05));
      opacity: 0;
      transition: opacity 0.3s ease;
    }

    .card:hover {
      transform: translateY(-6px);
      box-shadow: var(--shadow);
    }

    .card:hover::before {
      opacity: 1;
    }

    .card-preview {
      position: relative;
      padding: 32px;
      background: rgba(255, 255, 255, 0.02);
    }

    .card-icon {
      font-size: 4rem;
      opacity: 0.6;
      filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.3));
      transition: transform 0.3s ease;
    }

    .card:hover .card-icon {
      transform: scale(1.1) rotate(5deg);
    }

    .card-content {
      position: relative;
      padding: 24px;
      z-index: 1;
    }

    .card-title {
      font-size: 1.25rem;
      font-weight: 700;
      margin-bottom: 12px;
      color: var(--text);
    }

    .card-description {
      font-size: 0.9375rem;
      color: var(--text-muted);
      line-height: 1.6;
      margin-bottom: 16px;
    }

    .card-tags {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
    }

    .tag {
      padding: 4px 12px;
      background: rgba(78, 161, 255, 0.1);
      border: 1px solid rgba(78, 161, 255, 0.2);
      border-radius: 6px;
      font-size: 0.75rem;
      color: var(--accent-1);
      font-weight: 600;
    }

    .tag.physics {
      background: rgba(168, 85, 247, 0.1);
      border-color: rgba(168, 85, 247, 0.2);
      color: var(--accent-2);
    }

    .tag.animation {
      background: rgba(236, 72, 153, 0.1);
      border-color: rgba(236, 72, 153, 0.2);
      color: var(--accent-3);
    }

    .tag.interactive {
      background: rgba(245, 158, 11, 0.1);
      border-color: rgba(245, 158, 11, 0.2);
      color: var(--accent-4);
    }

    .tag.game {
      background: rgba(56, 189, 248, 0.1);
      border-color: rgba(56, 189, 248, 0.2);
      color: var(--accent-5);
    }

    /* Footer */
    footer {
      text-align: center;
      padding: 40px 24px;
      color: var(--text-muted);
      font-size: 0.875rem;
      border-top: 1px solid var(--card-border);
      margin-top: 60px;
    }

    /* Animations */
    @keyframes fadeInUp {
      from {
        opacity: 0;
        transform: translateY(30px);
      }
      to {
        opacity: 1;
        transform: translateY(0);
      }
    }

    @keyframes pulse {
      0%, 100% { opacity: 1; }
      50% { opacity: 0.5; }
    }

    /* Stagger animation delays */
    .card:nth-child(1) { animation-delay: 0s; }
    .card:nth-child(2) { animation-delay: 0.05s; }
    .card:nth-child(3) { animation-delay: 0.1s; }
    .card:nth-child(4) { animation-delay: 0.15s; }
    .card:nth-child(5) { animation-delay: 0.2s; }
    .card:nth-child(6) { animation-delay: 0.25s; }
    .card:nth-child(7) { animation-delay: 0.3s; }
    .card:nth-child(8) { animation-delay: 0.35s; }
    .card:nth-child(9) { animation-delay: 0.4s; }
  </style>
</head>
<body>
  <!-- Auto-generated via tools/sync_webfront_index.py -->
  <canvas class=\"bg-canvas\" id=\"bgCanvas\"></canvas>

  <div class=\"container\">
    <header>
      <h1>Webfront</h1>
      <p class=\"subtitle\">探索交互式前端演示 · 物理模拟 · 视觉特效 · 创意动画</p>
      <div class=\"badge\">
        <span>纯前端实现</span>
        <span>·</span>
        <span>无需构建工具</span>
      </div>
    </header>

{{CARDS}}

    <footer>
      <p>CodingToys · Webfront · 保持简单（KISS）· 开箱即用</p>
    </footer>
  </div>

  <script>
    // Animated background with floating particles
    const canvas = document.getElementById('bgCanvas');
    const ctx = canvas.getContext('2d');

    let width = canvas.width = window.innerWidth;
    let height = canvas.height = window.innerHeight;

    window.addEventListener('resize', () => {
      width = canvas.width = window.innerWidth;
      height = canvas.height = window.innerHeight;
    });

    const particles = [];
    const particleCount = 80;

    class Particle {
      constructor() {
        this.reset();
        this.y = Math.random() * height;
      }

      reset() {
        this.x = Math.random() * width;
        this.y = 0;
        this.vy = 0.3 + Math.random() * 0.5;
        this.vx = (Math.random() - 0.5) * 0.3;
        this.size = 1 + Math.random() * 2;
        this.opacity = 0.2 + Math.random() * 0.3;
        this.hue = 180 + Math.random() * 60; // Blue-ish tones
      }

      update() {
        this.x += this.vx;
        this.y += this.vy;

        if (this.y > height || this.x < 0 || this.x > width) {
          this.reset();
        }
      }

      draw() {
        ctx.beginPath();
        ctx.arc(this.x, this.y, this.size, 0, Math.PI * 2);
        ctx.fillStyle = `hsla(${this.hue}, 100%, 70%, ${this.opacity})`;
        ctx.fill();
      }
    }

    for (let i = 0; i < particleCount; i++) {
      particles.push(new Particle());
    }

    function animate() {
      ctx.clearRect(0, 0, width, height);

      particles.forEach((particle) => {
        particle.update();
        particle.draw();
      });

      requestAnimationFrame(animate);
    }

    animate();
  </script>
</body>
</html>
"""

def main() -> None:
    modules = load_metadata()
    cards_html = build_cards(modules)
    html = HTML_TEMPLATE.replace("{{CARDS}}", cards_html)
    INDEX_PATH.write_text(html + "\n", encoding="utf-8")
    print(f"Updated {INDEX_PATH.relative_to(ROOT)} with {len(modules)} modules.")

if __name__ == "__main__":
    main()
