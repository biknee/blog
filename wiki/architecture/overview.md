# 项目总览

## 基本信息

- **项目名**：个人博客（blog）
- **语言**：Go 1.26
- **模块名**：`blog`
- **端口**：:8080
- **用途**：本地个人博客，不对外部署

## 技术栈

| 组件 | 库 | 用途 |
|------|-----|------|
| HTTP 框架 | `gin-gonic/gin` | 路由、请求处理 |
| 模板引擎 | `html/template` | 服务端渲染 |
| Markdown | `yuin/goldmark` + `goldmark-highlighting` | 文章渲染 + 代码高亮 |
| 代码高亮样式 | `chroma` — monokai 主题 | 代码块着色 |
| YAML | `gopkg.in/yaml.v3` | Frontmatter 解析与序列化 |

## 目录结构

```
├── main.go                  # 入口：初始化 Service/Handler，注册路由，启动服务
├── go.mod / go.sum
├── handlers/
│   └── post.go              # HTTP 层：解析请求参数 → 调用 Service → 渲染模板
├── models/
│   └── post.go              # Post 结构体、ParseFile/ParseContent、Slugify
├── services/
│   └── post.go              # 业务层：文件 I/O（读/写/删 .md 文件）、列表查询
├── templates/
│   ├── header.html          # 公共页头（被 layout 引用或独立 include）
│   ├── footer.html          # 公共页脚
│   ├── home.html            # 文章列表页（分页卡片、标签筛选）
│   ├── post.html            # 文章详情页（Markdown 渲染后内容）
│   ├── new.html             # 新建文章表单
│   ├── edit.html            # 编辑文章表单
│   ├── 404.html             # 404 页面
│   └── md-ref.html          # Markdown 语法参考
├── static/css/
│   └── style.css            # Neo-Brutalism 样式
├── data/                    # 博客文章存储目录（.md 文件）
├── wiki/                    # 项目知识库（本目录）
└── CLAUDE.md                # Claude Code 项目指令
```

## 数据流

```
浏览器请求
  → main.go（路由匹配）
    → handlers/post.go（解析参数、调用 service）
      → services/post.go（读写 data/*.md 文件、返回 models.Post）
    → handlers/post.go（渲染 Goldmark → HTML）
  → templates/*.html（填入数据、输出 HTML）
浏览器渲染（Neo-Brutalism 样式）
```

## 关键入口

- **启动**：`go run main.go`
- **路由注册**：`main.go:26-33`
- **模板函数**：`main.go:19-22`（`add`、`subtract`）
- **静态文件**：`/static` → `./static` 目录
