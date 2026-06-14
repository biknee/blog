---
date: "2026-06-14"
slug: typescript-zod
tags:
    - TypeScript，Zod
title: TypeScript Zod
---

# Zod基础
弥补TS没有办法在运行时做检查的短板。通过单一数据源来约束并且通过语义化的方式来约束,类似于一个运行时的TS。
```
//import * as z from "zod";
import { z } from "zod";

const User = z.object(
  name: z.string().describe("姓名"),
  age: z.number().describe("年龄"),
  hi: z.coerce.number().describe(""), // 宽容模式，自动纠错
  start: z.iso.date(),
  end: z.iso.date(),
  sTime: z.string(), // 假设格式是 00:00:00.000
).
refine(data => { // 约束start<end
  return new Date(data.start) < new Date(data.end);
}).
refine(data => {
  return data.sTime.split(":").lenth === 3 && data.start.includes(".");
}).
describe("用户的基础协议");
type U = z.infer<typeof User>;
校验
cosnt re = User.parse({name:'awd',age:23}); // safeParse 安全校验
console.log(re);

```