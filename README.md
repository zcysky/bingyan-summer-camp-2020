# 冰岩作坊程序组2020夏令营

***欢迎参加冰岩作坊夏令营！***



## 前言

- 请先Fork(右上角)此仓库，本次夏令营要求代码、日报等全部托管在你们fork后的github仓库中
- 日报和周报等不需要写太多，只需要介绍每天学习了什么，以及适当记录你认为的重点即可
- 之后的代码检查(code review)，采用pull request(PR)的形式

## 操作说明

- fork 此仓库
- 在你的仓库操作的时候请不要对他人目录进行任何操作
- 你的操作权限仅限于你的目录，目录名字为你的 github ID，若仓库中没有你的目录请自行创建
- 提交 PR 的时候自行查看是否存在代码冲突，如果存在自行解决之后再提交 PR
- 提交 PR 是提交到 dev 分支，不是 master 分支
- 目录结构推荐如下：
  - reports文件夹 - 日报
  - source文件夹 - 源码，各项目创建不同的文件夹

## 学习安排

### 1.语言

- Golang语法
  - [官方链接](https://golang.org/)
  - [官方中文教程](https://tour.go-zh.org/welcome/1)
  - [语言规范](https://go-zh.org/ref/spec)

- 书籍推荐
  - [《The Go Programming Language》中文版](https://www.gitbook.com/book/yar999/gopl-zh/details)
  - [《Effective Go》中英双语版](https://www.gitbook.com/book/bingohuang/effective-go-zh-en/details)
  - [Go语言实战](http://download.csdn.net/download/truthurt/9858317)
  - [Go Web编程](https://wizardforcel.gitbooks.io/build-web-application-with-golang/content/index.html)可以了解基本web开发，比较推荐入门

### 2.框架

> 在学习框架的过程中，了解一下MVC架构，并在热身项目中加以应用。推荐gin和echo二选一

- gin

  - [gin英文文档](https://github.com/gin-gonic/gin)
  - [ Gin 文档中文翻译](https://learnku.com/docs/gin-gonic/2018/gin-readme/3819)
- echo
  - [echo英文文档](https://echo.labstack.com/guide)
  - [echo文档中文翻译](http://go-echo.org/)
- beego

  - [beego: 简约 & 强大并存的 Go 应用框架](https://beego.me/docs/intro/)

- Iris

  - [Iris英文文档](https://github.com/kataras/iris)

  - [Iris文档中文翻译](https://studyiris.com/doc/)

- 其他

### 3. HTTP相关

- HTTP请求方法：GET、POST、PUT、UPDATE等

- HTTP状态码：404、200、400、401、301、500等

- HTTP数据传输格式：[json](https://www.runoob.com/json/json-syntax.html)、form表单

- HTTP报文格式（大致了解就行、不用深入学习）

- 前后端如何交互？前后端分离是什么？

  前端如何获取后端返回的数据，如何发送请求，后端如何根据前端发过来的请求，回应请求，如何辨别不同的请求
  
  推荐阅读：《图解HTTP》

### 4. 数据库相关

- MySQL（推荐优先学习）
- MongoDB（后期推荐学习、可以在夏令营之后研究，有能力的可以夏令营用，和go搭配比较好用）
- Redis（基于内存的非关系型数据库）

### 5. 其他知识

**认证：**

熟悉以下三种前后端认证方式，一般在登录时使用

- cookie
- session
- JWT

**杂项**

- docker
- nginx

**加密算法：**

- 对称加密
- 非对称加密
- 哈希算法

### 6. 相关工具

- 编辑器：goland、vscode

- 后台接口测试工具：postman

  

## 阶段一：热身项目

**成员管理系统**

实现内容：

- 管理员和普通用户

- 用户注册和登录

  用户信息包括用户ID、密码（数据库中加密）、昵称、手机号、邮箱地址

- 管理员
  - 删除普通用户
  - 获取一个成员、所有成员信息

- 普通用户
  
  - 更改个人信息





## 阶段二：QQ机器人：

> 先做能做的，不必按顺序做。

#### 第一层次

运行QQ机器人。此处推荐使用CoolQ HTTP API或者Mirai. 设计一个简单的web后端，使得你的机器人能够对传入的消息进行应答。

你可以选择在本地运行CoolQ HTTP API 插件，当然更推荐你在云服务器上完成配置。

（如何获得免费的云服务器？https://developer.aliyun.com/adc/student/?spm=5176.12901015.0.i12901015.6462525crmOfkf）



#### 第二层次

你的QQ机器人加入了备忘录功能。具体来说分为：

- 添加事件
- 删除事件
- 查询所有已添加事件

你应该采取某种方式对添加的事件进行永久性的存储。

*进阶：采用数据库进行存储，并且可以对不同账号单独存储其事件。*



**完成到该层次内容，可以视为阶段二完成，当然有能力的同学可以挑战三层次内容**

#### 第三层次

你的QQ机器人在此基础上添加了提醒功能，具体来说：

- 每个事件具有一个DDL，在DDL到来之前会对你进行提醒。

**该层次的进阶内容**

- 采用数据库进行存储，可以对不同账号单独存储其事件。
- 用户可以为每个不同任务自定义提醒间隔和提早时间，例如我可以指定对于某事件 “提前x小时，每隔y分钟提醒一次。”



## 三阶段准备！

在完成第二阶段的情况下，视完成度发布第三阶段任务。





## 项目部署

### 1. 配置nginx

学习配置 nginx 做中间代理层，具体可从以下链接中选取部分学习，作为示例，夏令营之后可以好好研究，当然夏令营期间有时间也可以自行研究，遇到坑可以问我们。

[nginx 配置简介](https://juejin.im/post/5ad96864f265da0b8f62188f)

[openresty 实践](https://juejin.im/post/5aae659c6fb9a028d375308b)

### 2. 配置 docker

[Docker 从入门到实践](https://yeasy.gitbooks.io/docker_practice/content/install/ubuntu.html)

[Docker 实践](https://juejin.im/post/5b34f0ac51882574ec30afce)

### 3. 配置域名https (不要求)

前提：有已经备案的域名，有服务器

[Let's Encrypt 给网站加 HTTPS 完全指南](https://ksmx.me/letsencrypt-ssl-https/?utm_source=v2ex&utm_medium=forum&utm_campaign=20160529)