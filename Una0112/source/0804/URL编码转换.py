# -*- coding: utf-8 -*-
"""
Created on Mon Aug  3 16:15:22 2020

@author: 16038
"""

import urllib.parse
data= {"kw":"贴吧"}
# 通过 urlencode() 方法，将字典键值对按URL编码转换，从而能被web服务器接受。
data = urllib.parse.urlencode(data)
print(data)  # kw=%E8%B4%B4%E5%90%A7
# 通过 unquote() 方法，把 URL编码字符串，转换回原先字符串。
data = urllib.parse.unquote(data)
print(data)  # kw=贴吧