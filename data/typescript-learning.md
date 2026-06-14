---
date: "2026-06-14"
slug: typescript-learning
tags:
    - TypeScript
title: TypeScript Learning
---

# 项目初始化
npm init -y
package.json
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
  "type": "commonjs",
  "dependencies":{},
  "devDependencies":{}
}

#### package.json
scripts:
- build:"tsc ." 构建产物 “build”
- dev:"tsc --watch" 监听的方式启动，内容被修改会自动更新启动

devDependencies：
- typescript:"5.9.3" 指定TS版本
- @types/node:"25.9.3" 指定node版本，安装好后需在tsconfig的types中指定

## 插件安装
1. 需要明确：插件是“dependencies”，还是“devDependencies”，版本是多少
2. 确定好后：pnpm i (pnpm install缩写)安装插件；没有pnpm：npm install -g pnpm
3. tsc介绍: TS内置编译器，TS代码不能直接执行，必须通过编译成js文件
4. 配置文件：npx tsc --init 得到tsconfig.json，每个项目必须

#### tsconfig.json
- module: "nodenext" 如果配置了moduleResolution则修改为esnext，如果是针对nodejs则不改
- outDir:"./dist" tsc编译文件输出位置
- types:["node"] 补充依赖
- sourceMap:true 输出结果需不需要sourceMap
- declaration:true 输出结果需不需要类型声明文件
- declarationMap:true 输出结果需不需要类型声明文件的map文件
**以下推荐补全**
- strict: true
- jsx: "react-jsx", 针对jsx直接指定成react-jsx
- //verbatimModuleSyntax:true 注释掉
- isolatedModules:true
- noUncheckedSideEffectImports:true
- moduleDetection:"force"
- moduleResolution:"bundler" 一般情况下前端应用直接指定成bundler，同时修改module，node则不加
- skipLibCheck: true


# 进阶
- 类型定义
  - 基础类型
  - 引用类型
  - 自定义类型
- 类型定义更灵活：泛型
- 函数灵活定义：泛型
- 两个兜底符号：?可选属性/!非空断言

### 声明变量的时候，需要确定类型
基础类型：
- string 字符串
const modelName: string = 'qwen3:7b';
- number 数字
const temperature: number = 0.7;
- boolean 布尔
const isStreaming: boolean = true;
- undefined
- null
- symbol

引用类型：
- string[] 数组
string | number | boolean 联合类型，表示可以是其中之一
const messages: {type: string | number | boolean; text: string}[] = [{type:123,text:"avs"}];
const messages: {type: "ai"| "user" | "system"; text: string}[] = [{type:"ai",text:"avs"}];
- Array<string> 数组：泛型写法
- [number,number,number] 元组：固定长度，固定类型（如向量）
const embeddingVector: [number,number,number] = [0.1,0.23,0.32]

其他：
- any 放弃类型检查，不建议使用
- unknown 不知道是什么类型，使用前必须做类型检查

### 接口vs类型别名：
interface User{
  name: string;
  age: number;
}

type User = {
  name: string;
  age: number;
}
主要区别：拓展的方式不一样
interface NewUser extends User {}
type NewUser = User & {hobby: string[]} 交叉类型

### 函数声明：
function invoke(prompt: string | {type: "ai"| "user" | "system"; text: string}[]): string {
  // 类型保护的方式来访问
  if(typeof prompt === 'boject'){ // 类型保护，类型收敛
    prompt[0].
  } else {
    
  }
  return 'str';
}
泛型声明：
function invoke<T>(prompt: T): string{
  // 类型保护的方式来访问
  if(typeof prompt === 'boject'){ // 类型保护，类型收敛
    prompt[0].
  } else if (typeof prompt === 'string'){
    
  } else {
  
  }
  return 'str';
}

### 类型工具方法：Omit：去掉type中的某个属性 Pick：摘取type中的某个属性
const tinyUser: Omit<User, "age"> = {
  name: "afc"
};

### 联合消息类型：
interface AIMsg {role: "assistant"; tool_calls?:any[]; content: string | null}
interface ToolMsg {role: "tool"; tool_call_id: string; content: string}
interface UserMsg {role: "user"; content: string}
interface ChatMsg = AIMsg | ToolMsg | UserMsg;
const message: ChatMsg = {role: 'user', content: ""}

### 类型推导：
const aaa = {name: 'a', age: 13, hobby:["cs", "ball"]}
type A = typeof // aaa 声明一个对象是aaa类型
const test: A = {name:'', age: 12, hobby:[]}