# CLAUDE.md

本文件为 Claude Code 在本仓库中工作时提供指引。

## 项目：个人博客（本地）

一个仅限本地使用的个人博客网站，使用 Golang + Gin 构建，博客文章以 Markdown 文件（YAML frontmatter）形式存储。

## 角色定位

你是该项目的**开发者**。负责实现功能、修复 Bug、编写简洁地道的 Go 代码。不要做半成品实现。优先编辑已有文件，而非创建新文件。除非 WHY 不够明显，否则不写注释。

## 开发铁律

**每完成一个接口或模块，必须先测试通过，确认功能正常后才能继续下一个。**
- 接口完成后立即用 `curl` 验证 HTTP 状态码和响应内容
- 模块完成后写一个最小验证脚本或直接调用确认
- 测试不通过绝不往下走

## 技术栈

- **语言**：Go（最新稳定版，当前 1.26.2）
- **HTTP 框架**：gin-gonic/gin
- **Markdown 渲染**：yuin/goldmark + goldmark-highlighting
- **Frontmatter 解析**：gopkg.in/yaml.v3
- **模板引擎**：Go html/template（服务端渲染，非 SPA）
- **端口**：8080

## 项目结构

```
claudeCode/
├── main.go                  # 入口，路由注册
├── go.mod / go.sum
├── handlers/
│   └── post.go              # HTTP handlers（薄层，委托给 services）
├── models/
│   └── post.go              # Post 结构体，YAML frontmatter 解析
├── services/
│   └── post.go              # 业务逻辑：Markdown 文件 CRUD
├── templates/
│   ├── layout.html          # 基础布局（页头、页脚、Neo-Brutalism 外壳）
│   ├── home.html            # 博客列表（分页）
│   ├── post.html            # 单篇文章（Markdown 渲染后）
│   ├── new.html             # 新建文章表单
│   └── edit.html            # 编辑文章表单
├── static/
│   └── css/
│       └── style.css        # Neo-Brutalism 样式
├── data/                    # 博客文章 .md 文件
├── wiki/                    # 项目知识库（架构、功能、约定、错误记录）
└── CLAUDE.md
```

## 设计风格：Neo-Brutalism

- 配色：白底 `#FFF`，黑字/黑边框 `#000`，黄色主色 `#FFBE0B`，粉色强调 `#FF006E`
- 圆角：全部 **0px**
- 边框：卡片、按钮、输入框统一 **3px solid black**
- 阴影：硬投影 `5px 5px 0 #000`（无模糊）
- 字体：标题用 Space Grotesk（粗体大写），代码/标签用 JetBrains Mono，正文用系统无衬线字体
- 按钮：实心底色 + 3px 黑边框 + 硬阴影；`:active` 时 `translateY(3px)`（按下感）
- 禁用渐变、禁用 opacity 淡入淡出——位移是唯一的动效手段
- 标签/徽章：黄底黑字，2px 黑边框，小硬阴影

## 博客文件格式

`data/` 下的每个 `.md` 文件使用 YAML frontmatter：

```markdown
---
title: "文章标题"
date: "2026-05-29"
tags: ["Go", "Blog"]
slug: "article-slug"
summary: "可选摘要"
---

用 **Markdown** 写的正文...
```

- `slug` 缺失时从标题自动生成（小写 + 连字符）
- 文件名格式为 `<slug>.md`，存放于 `data/` 目录
- frontmatter 中的 `slug` 字段优先级高于文件名

## API 路由

| 方法 | 路径 | 功能 |
|------|------|------|
| GET | `/` | 博客列表（分页，每页 10 篇） |
| GET | `/post/:slug` | 查看文章 |
| GET | `/new` | 新建文章表单 |
| POST | `/post` | 创建文章 |
| GET | `/edit/:slug` | 编辑文章表单 |
| POST | `/post/:slug` | 更新文章 |
| POST | `/post/:slug/delete` | 删除文章 |
| GET | `/tag/:tag` | 按标签筛选 |

## 编码规范

- Handler 保持薄：解析请求 → 调用 service → 渲染模板
- Service 包含所有文件 I/O 和业务逻辑
- Model 定义数据结构和序列化
- 错误处理：文章不存在返回 404，文件错误返回 500，输入非法返回 400
- 无需认证、无需 CSRF、无需限流——本地专用
- 使用 `log.Printf` 打日志，`log.Fatalf` 仅限 main.go 中不可恢复的启动错误

## 知识库（wiki/）

`wiki/` 目录是项目的结构化知识库，**每次对话开始时应先浏览其目录结构**，了解已有内容。其目的是节约 Token、避免重复扫描和重复犯错。

### 使用时机

| 场景 | 动作 |
|------|------|
| 需要了解项目结构 / 定位代码 | **先查** `wiki/architecture/`，查不到再读源码 |
| 新增功能 | **先查** `wiki/features/inventory.md`，确认是否已存在或冲突 |
| 遇到报错 | **先查** `wiki/errors/known-errors.md`，确认是否曾犯过 |
| 修改样式或模板 | **先查** `wiki/conventions/go-coding.md` 的设计风格部分 |
| 不确定编码风格 | **先查** `wiki/conventions/` |

### 维护规则

- 涉及知识库覆盖内容的代码变更，**同步更新**对应知识文件
- 旧知识完全过时时**删除**，避免误导
- 同类新知识以**追加**形式写入已有文件，不新建碎片文件
- 新增错误记录时复制 `wiki/errors/known-errors.md` 中的模板格式

### 目录结构

```
wiki/
├── README.md                  # 知识库用途与规则
├── architecture/              # 架构文档（项目总览、路由映射）
├── features/                  # 功能清单（已实现功能）
├── conventions/               # 编码与设计约定
└── errors/                    # 已知错误与修复记录
```

## 环境

- **操作系统**：Windows 11
- **Shell**：Bash (Git Bash)
- **编辑器**：任意——文件通过 Claude Code 工具编辑
