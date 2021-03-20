# 爱课程API

## 课程
### 课程数据库
```
{
    "lid" 课程编号 int
    "course" 课程名称 string
    "type" 课程类别 string
    "credit" 学分 int
    "target" 针对对象 string
    "arrange" 时地安排
    	[	
        	{
                "time" 开课时间 string
                "location" 开课地点 string
			}
            ,...
        ]
    "evaluation" 课程评价
    	{
        	"good"
            "soso"
            "bad"
        }
    "tag":
        [
            {
                "type":string
                "value":string
            }
            ,...
        ]
    
}
```
### 所有课程概览
请求
```
GET  /allcourses?type= &time= &location= &atdc= &exam=
0不选择
页面第一次加载会发一次全0请求

type 类别
	沟通与管理
    科技与环境
    历史与文化
    文学与艺术
    社会与经济
    思维与方法
time 时间
position 地点
	east
    west
atdc 点名方式
	点名
    签到
    不点名不签到
exam 考核方式
	论文
    考试
    其他
```
返回
```json
[
    {
        "lid":int
        "course": string
        "type":string
        "teacher":string
        "tag":
        [
            {
                "type":string
                "value":string
            }
            ,...
        ]
    }
    ,...
]
```
### 特定课程详情
请求
```
GET /course/:lid
lid 课程编号
```
返回
```json
{
    "lid":int
    "course": string
    "type":string
    "credit":int
    "teacher":string
    "tag":
        [
            {
                "type":string
                "value":string
            }
            ,...
        ]
    "evaluation" 
        {
            "good":int 
            "soso":int
            "bad":int
        }

    "comment":?
}
```
## 评论
### 评论数据库
```
{
	"cid":int
    "uid":int
    "lid":int
    "user":string
    "icnt":boolean  匿名
    "term-end":boolean 期末情报
    "avatar":string  (头像url)
    "time":int     (unix timestamp)
    "content":string
    "img":string
    "like":int
    "subcmt-num":int
    "subcmt":[
    	{
        	"uid":int
            "user":string
			"icnt":boolean  
            "avatar":string
            "time":int
            "content":string
        }
        ,...
    ]
}
```
### 所有评论概览
请求
```
GET /allcomments/:lid
```
返回<br>
匿名用户对名字和头像进行处理后返回

原小程序中子评论无法匿名，此处评论添加匿名功能
```
[
    {
        "user":string
        "icnt":boolean  匿名
        "term-end":boolean 期末情报
        "avatar":string  (头像url)
        "time":int     (unix timestamp)
        "content":string
        "like":int
        "subcmt-num":int
        ]
    }
    ,...
]
```

### 特定评论详情
请求
```
GET /comment/:cid
```
返回
```

{
    "user":string
    "icnt":boolean  匿名
    "term-end":boolean 期末情报
    "avatar":string  (头像url)
    "time":int     (unix timestamp)
    "content":string
    "like":int
    "subcmt-num":int
    "subcmt":[
    	{
    		"cid":int
        	"uid":int
            "user":string
			"icnt":boolean  
            "avatar":string
            "time":int
            "content":string
        }
        ,...
    ]
}
```

### 发表评论
请求
```
POST /user/comment/:lid
body:{
	"uid":int
	"avatar":string
	"user":string
	"time":int
	"atdc":string
    "exam":string
    "icnt":boolean
    "term-end":boolean
    "evaluation":string
    "content":string
    "img":string
}
```
### 发表子评论
请求
```
POST /user/subcmt/:cid
body:{
	"uid":int
	"icnt":bool
	"user":string
	"avatar":string
    "time":int
    "content":string
}
```
## 用户
### 用户数据集库
```
{
	"uid":int
    "openid":string
    "user":string
    "type":string    ("admin" 或 "general"）
    "avatar":string
}
```
### 登录

//使用RSA加密，登录前要向后端申请公钥，code要

请求
```
GET /token?code=
code为wx.login()返回code
```
使用auth.code2Session接口获得openid，检查该用户是否存在<br>
返回

```
jwt token
pay load 中含有用户uid，nickname，type
```
### 初始化
服务端无法直接获得用户信息，用户在登录后小程序发出更新用户信息请求<br>

请求

```
POST /user/:uid
body={
	"user":string
    "avatar":string
}
```
### 获取用户发表的所有评论

```
GET /user/:uid/allcomments
```

