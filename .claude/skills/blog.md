---
name: blog
description: 个人博客项目开发——基于 Markdown 文件的博客 CRUD 操作，Neo-Brutalism 设计风格
metadata:
  type: project
---

# Blog 项目技能

本技能为个人博客项目提供开发上下文——一个使用 Golang + Gin 的本地博客，Markdown 文件存储，Neo-Brutalism 设计。

## 开发铁律

**每完成一个接口或模块，必须先测试通过，确认功能正常后才能继续下一个。**
- 接口完成后立即用 `curl` 验证 HTTP 状态码和响应内容
- 测试不通过绝不往下走

## 速查

- **技术栈**：Go 1.26 + Gin + Goldmark + html/template
- **端口**：8080
- **数据**：`./data/*.md`（YAML frontmatter + Markdown 正文）
- **风格**：Neo-Brutalism（0px 圆角、3px 黑边框、硬阴影、黄色/粉色强调色）

## 常用操作

### 添加新功能
1. 如需新数据字段，先改 `models/post.go`
2. 在 `services/post.go` 中新增 service 方法
3. 在 `handlers/post.go` 中新增 handler
4. 在 `main.go` 中注册路由
5. 更新模板（如需要）
6. **立即用 curl 测试该接口，确认通过后再继续**

### 启动服务器
```bash
cd /c/C-Study/stuff/claudeCode && go run main.go
```
打开 http://localhost:8080

### 验证编译
```bash
cd /c/C-Study/stuff/claudeCode && go build ./...
```

### 测试接口
```bash
# 首页
curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/

# 查看文章
curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/post/hello-world

# 新建页面
curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/new
```

### 添加测试文章
在 `data/` 下创建一个带正确 YAML frontmatter 的 `.md` 文件（需含 title、date、tags、slug）。

## 架构说明

- **handlers** 是薄层——解析请求 → 调用 service → 渲染模板
- **services** 包含所有文件 I/O——读写 `data/*.md`
- **models** 定义 `Post` 结构体和 frontmatter 的序列化/反序列化
- **templates** 使用 `LoadHTMLGlob` 从磁盘加载，修改后无需重新编译即可生效
- 无数据库、无认证、无外部依赖（仅 Gin + Goldmark + YAML）
