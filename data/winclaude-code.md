---
date: "2026-05-29"
slug: winclaude-code
tags:
    - claude code，安装
title: win安装claude code
---

一、配置powershell代理：
1.安装最新的 PowerShell 7
iex "& { $(irm https://aka.ms/install-powershell.ps1) } -UseMSI"

2.在 PowerShell 中运行以下命令，这会自动用记事本打开（如果没有则创建）你的配置文件：
if (!(Test-Path $PROFILE)) { New-Item -Type File -Path $PROFILE -Force }; notepad $PROFILE

3.记事本打开后，将以下代码完整地复制并粘贴进去：
```
# 开启代理函数
function proxyOn {
    $env:HTTP_PROXY = "http://127.0.0.1:7897"
    $env:HTTPS_PROXY = "http://127.0.0.1:7897"
    
    Set-ItemProperty -Path "HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings" -Name ProxyEnable -Value 1
    Set-ItemProperty -Path "HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings" -Name ProxyServer -Value "127.0.0.1:7897"
    
    Write-Host " [√] Proxy Enabled" -ForegroundColor Green
}

# 关闭代理函数
function proxyOff {
    $env:HTTP_PROXY = $null
    $env:HTTPS_PROXY = $null
    
    Set-ItemProperty -Path "HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings" -Name ProxyEnable -Value 0
    
    Write-Host " [X] Proxy Disabled" -ForegroundColor Red
}

# 查看代理状态函数
function proxyState {
    Write-Host "--- Proxy Status Check ---" -ForegroundColor Cyan
    
    # 1. 检查 PowerShell 环境变量
    if ($env:HTTP_PROXY -or $env:HTTPS_PROXY) {
        Write-Host " [√] PowerShell Environment Proxy: ENABLED" -ForegroundColor Green
        Write-Host "     Address: $env:HTTP_PROXY" -ForegroundColor Gray
    } else {
        Write-Host " [X] PowerShell Environment Proxy: DISABLED" -ForegroundColor Red
    }
    
    # 2. 检查 Windows 系统全局注册表
    $regProxy = Get-ItemProperty -Path "HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings"
    if ($regProxy.ProxyEnable -eq 1) {
        Write-Host " [√] Windows System Global Proxy:  ENABLED" -ForegroundColor Green
        Write-Host "     Address: $($regProxy.ProxyServer)" -ForegroundColor Gray
    } else {
        Write-Host " [X] Windows System Global Proxy:  DISABLED" -ForegroundColor Red
    }
    Write-Host "--------------------------" -ForegroundColor Cyan
}
```
4.保存并使配置生效:
. $PROFILE

二、claude code安装
官网：https://claude.com/product/claude-code
安装：irm https://claude.ai/install.ps1 | iex
注：使用ccSwitch中转claude请求到ds需注意最新版本的claude格式不支持，需使用旧版claude code

旧版安装：
1.安装 Node.js 环境:官网安装包：https://nodejs.org/

2.安装完成后，必须彻底关闭当前的 PowerShell 7 窗口，重新打开一个新的（让系统加载 npm 环境变量）

3.使用 npm 降级安装真正稳定的旧版本：npm install -g @anthropic-ai/claude-code@2.1.148

4.验证claude --version