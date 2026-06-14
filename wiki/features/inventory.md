# 功能列表

## 文章管理

| 功能 | 状态 | 入口 | 说明 |
|------|------|------|------|
| 文章列表 | 已实现 | `GET /` | 按日期倒序，每页 12 篇，支持分页导航 |
| 查看文章 | 已实现 | `GET /post/:slug` | Markdown 渲染，代码语法高亮（monokai） |
| 新建文章 | 已实现 | `GET /new` → `POST /post` | 表单输入标题、标签（逗号分隔）、Markdown 正文 |
| 编辑文章 | 已实现 | `GET /edit/:slug` → `POST /post/:slug` | 加载已有内容到表单，保存覆盖 |
| 暂存草稿 | 已实现 | `POST /post/:slug/save-draft` | 编辑页暂存当前内容，不校验必填，留在编辑页 |
| 删除文章 | 已实现 | `POST /post/:slug/delete` | 直接删除 .md 文件，重定向到首页 |
| 按标签筛选 | 已实现 | `GET /tag/:tag` | 大小写不敏感匹配 |
| 自动生成 slug | 已实现 | 模型层 | 标题→小写+连字符，冲突时追加时间戳后缀 |
| 自动摘要 | 已实现 | 模型层 | frontmatter 无 summary 时取正文前 150 字符 |

## Markdown 渲染

| 功能 | 状态 | 说明 |
|------|------|------|
| 标准 Markdown | 已实现 | goldmark 渲染 |
| 代码语法高亮 | 已实现 | goldmark-highlighting + chroma monokai 主题 |

## 模板与前端

| 功能 | 状态 | 说明 |
|------|------|------|
| Neo-Brutalism 风格 | 已实现 | 0px 圆角、3px 黑边框、硬阴影、黄黑配色 |
| 响应式布局 | 未实现 | 仅桌面端优化 |
| Markdown 语法参考页 | 已实现 | `/md-ref`（无路由，静态展示） |

## 系统

| 功能 | 状态 | 说明 |
|------|------|------|
| 404 页面 | 已实现 | 文章不存在时渲染 404.html |
| PostToolUse hook | 已实现 | Edit/Write 后自动 `go build` 检查编译 |
