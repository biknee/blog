# Go 编码规范

## 分层约定

```
main.go         → 初始化依赖、注册路由、启动服务（仅此处可用 log.Fatalf）
handlers/       → 薄层：解析请求 → 调用 service → 渲染模板（不写业务逻辑）
services/       → 所有文件 I/O 和业务逻辑
models/         → 数据结构定义、序列化/反序列化
```

## 错误处理

- 文章不存在 → 404（`handlers/post.go:64-69`）
- 文件读写错误 → 500
- 输入非法（标题/内容为空）→ 400，返回表单页并显示错误
- `log.Fatalf` 仅限 `main.go` 中不可恢复的启动错误

## 命名

- 文件名：小写 + 连字符（如 `post.go`）
- 包名：小写单数（`handlers`、`services`、`models`）
- 导出符号：大写驼峰
- 私有符号：小写驼峰

## 依赖

- 不新增不必要的第三方依赖
- 优先使用标准库

## 安全

- 无需认证、无需 CSRF、无需限流——本地专用
- 用户输入（标题、标签、内容）需 `TrimSpace`
- 模板中 `Content` 字段使用 `template.HTML`（已由 goldmark 渲染，可信任）

## 设计风格（Neo-Brutalism）

| 属性 | 值 |
|------|-----|
| 背景色 | `#FFF` |
| 前景色 | `#000` |
| 主色 | `#FFBE0B`（黄） |
| 强调色 | `#FF006E`（粉） |
| 圆角 | 0px |
| 边框 | 3px solid black |
| 阴影 | `5px 5px 0 #000` |
| 字体-标题 | Space Grotesk（粗体大写） |
| 字体-代码 | JetBrains Mono |
| 字体-正文 | 系统无衬线 |
| 动效 | 仅 `translateY` 位移，禁用 opacity / 渐变 |
