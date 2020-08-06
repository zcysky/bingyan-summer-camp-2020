# -*- coding: utf-8 -*-
"""
Created on Mon Aug  3 16:15:22 2020

@author: 16038
"""

import urllib.request
import re
import ssl

# 取消代理验证
ssl._create_default_https_context = ssl._create_unverified_context


class Spider:
    def __init__(self):
        self.page = 188076151
        self.switch = True

    def loadPage(self):
        """
           作用：打开页面
        """
        url = "https://www.douban.com/group/topic/" + str(self.page) + "/"
        user_agent = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.157 Safari/537.36'
        headers = {'User-Agent': user_agent}
        req = urllib.request.Request(url, headers=headers)
        response = urllib.request.urlopen(req)
        html = response.read().decode('utf-8')
        pattern = re.compile(r'<div.*?class="content">(.*?)</div>', re.S)
        content_list = pattern.findall(html)
        self.dealPage(content_list)

    def dealPage(self, item_list):
        """
            @brief 处理得到的糗事列表
            @param item_list 得到的糗事列表
            @param page 处理第几页
        """
        for item in item_list:
            item = item.replace('<span>', "").replace('<span class="contentForAll">查看全文', "").replace("</span>","").replace("<br/>", "").replace("\n", "")
            self.writePage(item)

    def writePage(self, text):
        """
            @brief 将数据追加写进文件中
            @param text 文件内容
        """
        myFile = open("./shenzu.md", 'a') 
        myFile.write(text + "\n\n")
        myFile.close()

    def startWork(self):
        """
            控制爬虫运行
        """
        while self.switch:
            self.loadPage()
            command = input("如果继续爬取，请按回车（退出输入quit)")
            if command == "quit":
                self.switch = False
            self.page += 1
        print("爬取结束！")


if __name__ == '__main__':
    qiushiSpider = Spider()
    qiushiSpider.startWork()