---
date: "2026-06-14"
slug: typescript-learning
tags:
    - TypeScript
title: TypeScript Learning
---

# 项目初始化
npm init -y

{
  "name": "tslearn",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "keywords": [],
  "author": "",
  "license": "ISC",
  "type": "commonjs"
}

## 插件安装
1. 需要明确，插件是“dependencies”，还是“devDependencies”，版本是多少
2. 确定好后，pnpm i (pnpm install缩写)安装插件