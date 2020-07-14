# Daily Report

## 2020.7.13

### 学习内容

* Goland IDE使用
* 初步学习使用Go mod构建项目
* Go的基本语法：变量，函数，循环、选择、包，简单输入输出，`math.big`包的简单使用

### 重点

* Go环境的安装。由于命令行代理走snap不熟，snap国外源自动安装Go环境速度慢。而apt中的Golang版本为1.10未及时更新，最终选择使用官方文档推荐的下载`tar.gz`包并向`\etc\profile`文件中添加go的bin路径
* Go mod。最初未使用Go mod构建项目时每次创建新项目都需新建GOPATH参数。使用Go mod可免去此步骤。同时，Go mod个人感觉部分功能与Java中的Maven类似，可以自动添加依赖，而不用手动控制。
* 学会看官方文档。大数相加任务最初未仔细看官方文档，自己手写数字字符串的加法。后查`math.big`文档，发现有直接将字符串转为`big.Int`的函数。



## 2020.7.14

### 学习内容

* 发Http请求
* Go中对Json的处理
* go routine使用

### 重点

* Http请求的发送。使用client+NewRequest发送请求，返回Json
* Json在Go中是[]byte，对其进行解码，用struct与其对应
* go routine，对并发知识的理解。