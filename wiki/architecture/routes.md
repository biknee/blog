# 路由映射

## 路由表

| 方法 | 路径 | Handler 方法 | Service 方法 | 模板 | 功能 |
|------|------|-------------|-------------|------|------|
| GET | `/` | `Home` | `ListPosts(page, 12)` | `home.html` | 文章列表分页 |
| GET | `/post/:slug` | `View` | `GetPost(slug)` | `post.html` | 查看单篇文章 |
| GET | `/new` | `NewForm` | — | `new.html` | 新建文章表单 |
| POST | `/post` | `Create` | `CreatePost(title, tags, content)` | 重定向 | 创建文章 |
| GET | `/edit/:slug` | `EditForm` | `GetPost(slug)` | `edit.html` | 编辑文章表单 |
| POST | `/post/:slug` | `Update` | `UpdatePost(slug, title, tags, content)` | 重定向 | 更新文章 |
| POST | `/post/:slug/save-draft` | `SaveDraft` | `SaveDraft(slug, title, tags, content)` | 重定向 `/edit/:slug` | 暂存草稿 |
| POST | `/post/:slug/delete` | `Delete` | `DeletePost(slug)` | 重定向 `/` | 删除文章 |
| GET | `/tag/:tag` | `Tag` | `ListPostsByTag(tag)` | `home.html` | 按标签筛选 |

## Handler → Service 完整映射

```
Home         → ListPosts(page, perPage)       → ([]Post, total, error)
View         → GetPost(slug)                  → (*Post, error)
NewForm      → 无（直接渲染模板）
Create       → CreatePost(title, tags, cont)   → (*Post, error)
EditForm     → GetPost(slug)                  → (*Post, error)
Update       → UpdatePost(slug, title, t, c)  → (*Post, error)
SaveDraft    → SaveDraft(slug, title, t, c)  → (*Post, error)
Delete       → DeletePost(slug)               → error
Tag          → ListPostsByTag(tag)            → ([]Post, error)
```

## 代码位置

- 路由注册：`main.go:26-34`
- Handler 定义：`handlers/post.go:34-199`
- Service 定义：`services/post.go:24-199`
- 模板目录：`templates/`
