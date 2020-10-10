# -*- coding: utf-8 -*-
"""
Created on Mon Aug  3 16:15:22 2020

@author: 16038
"""

import urllib.request
import urllib.parse
import ssl

# 取消代理验证
ssl._create_default_https_context = ssl._create_unverified_context

# User-Agent
headers = {"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"}
# post 请求参数
data = {
    "type": "11",
    "interval_id": "100:90",
    "start": "0",
    "limit": "10"
}
data = urllib.parse.urlencode(data).encode("utf-8")

# url 作为Request()方法的参数，构造并返回一个Request对象
request = urllib.request.Request("https://movie.douban.com/j/chart/top_list?", data=data, headers=headers)

# Request对象作为urlopen()方法的参数，发送给服务器并接收响应
response = urllib.request.urlopen(request)

# 类文件对象支持 文件对象的操作方法，如read()方法读取文件全部内容，返回字符串
html = response.read().decode("utf-8")

# 打印字符串
print(html)