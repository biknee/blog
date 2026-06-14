# 错误记录

---

## 错误 1：JSON 配置文件包含原始控制字符

**日期**：2026-06-14

**现象**：`/doctor` 报告 `.claude/settings.local.json` 为无效 JSON

**根因**：hook command 字符串中嵌入了原始 ANSI 转义字符（`0x1b`，即 ESC）。JSON 规范不允许字面量控制字符出现在字符串中。

**定位**：`node -e "JSON.parse(...)"` 报 `Bad control character in string literal`，`xxd` 确认文件中存在 `0x1b` 字节。

**修复**：
```bash
esc=$(printf '\x1b')
sed -i "s/$esc/\\\\u001b/g" .claude/settings.local.json
```

**预防**：在 JSON 配置中使用 ANSI 颜色码时，使用 `` 转义形式而非原始 ESC 字符。

**教训**：JSON 字符串中禁止字面量控制字符（`\x00`–`\x1f`），需用 `\u00xx` 转义。在 JS 中通过 `String.fromCharCode(27)` 或 `\x1b` 匹配原始字节时，注意 shell 变量和 encoding 可能导致替换失败——用 `sed` 做二进制级替换更可靠。

---

## 模板

后续新增错误按此格式追加：

```
## 错误 N：简短标题

**日期**：YYYY-MM-DD

**现象**：描述表面看到的问题

**根因**：解释根本原因

**定位**：如何快速定位到问题所在

**修复**：具体修复命令或代码变更

**预防**：如何避免再次发生
```
