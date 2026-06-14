# 个人博客

基于 Go + Gin 的本地个人博客，Markdown 文件存储，服务端渲染，Neo-Brutalism 设计风格。

## 功能

- Markdown 撰写文章（YAML frontmatter），支持代码高亮
- 文章列表分页、按标签筛选、查看 / 新建 / 编辑 / 删除
- 独立的 Neo-Brutalism 主题——硬边框、硬阴影、高饱和配色，位移驱动动效

## 技术栈

| 领域 | 选型 |
|------|------|
| 语言 | Go 1.26 |
| HTTP 框架 | gin-gonic/gin |
| 模板引擎 | Go html/template |
| Markdown | yuin/goldmark + goldmark-highlighting |
| 代码高亮 | chroma (monokai 主题) |
| Frontmatter | gopkg.in/yaml.v3 |

## 目录结构

```
├── main.go                  # 入口，路由注册
├── go.mod / go.sum
├── handlers/
│   └── post.go              # HTTP handlers（薄层，解析请求 → 调用 service → 渲染模板）
├── models/
│   └── post.go              # Post 结构体，YAML frontmatter 解析
├── services/
│   └── post.go              # 业务逻辑：Markdown 文件 CRUD
├── templates/
│   ├── header.html          # 公共页头
│   ├── footer.html          # 公共页脚
│   ├── home.html            # 文章列表（分页）
│   ├── post.html            # 文章详情
│   ├── new.html             # 新建文章表单
│   ├── edit.html            # 编辑文章表单
│   ├── 404.html             # 404 页面
│   └── md-ref.html          # Markdown 语法参考
├── static/
│   └── css/
│       └── style.css        # Neo-Brutalism 样式
└── data/                    # 博客文章 .md 文件
```

## 快速开始

```bash
# 安装依赖
go mod tidy

# 启动（默认 :8080）
go run main.go
```

浏览器打开 `http://localhost:8080`。

## API 路由

| 方法 | 路径 | 功能 |
|------|------|------|
| GET | `/` | 文章列表（分页，每页 12 篇） |
| GET | `/post/:slug` | 查看文章 |
| GET | `/new` | 新建文章表单 |
| POST | `/post` | 创建文章 |
| GET | `/edit/:slug` | 编辑文章表单 |
| POST | `/post/:slug` | 更新文章 |
| POST | `/post/:slug/delete` | 删除文章 |
| GET | `/tag/:tag` | 按标签筛选 |

## 文章格式

`data/` 下的 `.md` 文件使用 YAML frontmatter：

```markdown
---
title: "文章标题"
date: "2026-06-14"
tags: ["Go", "Blog"]
slug: "hello-world"
summary: "可选摘要"
---

Markdown 正文。
```

- `slug` 缺失时从标题自动生成（小写 + 连字符）
- 文件名需与 slug 一致：`<slug>.md`
- frontmatter 中的 `slug` 优先级高于文件名

## 设计风格

Neo-Brutalism：白底黑字黑边框，`#FFBE0B` 黄主色 + `#FF006E` 粉强调色。全部 0px 圆角，3px 实心黑边框，硬投影 `5px 5px 0 #000`。按钮压下使用 `translateY` 位移而非 opacity 渐变。
