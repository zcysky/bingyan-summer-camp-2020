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

--------------

阶段一热身项目：
\- 学习 MVC 模式和 RESTful 设计规范
\- 撰写 API 文档，撰写完成后请与导师交流再进行下一步
\- 设计 MVC 模型，请与导师交流再进行下一步
\- 开发
\- 测试（Postman）
\- 在远端运行



该项目需求参见https://github.com/zcysky/bingyan-summer-camp-2020，注意事项：
\- 用户身份验证使用 Authorization 头，Bearer 后接 JWT 令牌的模式
\- 用户注册需要进行邮箱验证，请自行寻找发送邮件的方法并向注册邮箱发送随机验证码进行确认
\- 上述邮箱验证码请保存在 Redis 中（github.com/go-redis/redis)，用后或过期记得删除，其他信息请保存在 MongoDB 中
\- 配置文件保存在 config.json 中

------------

### Restful

```
# register
POST /user
{
	"name": str
	'password' str
	"email": str
	"phone": str
	"register_code": str
}

# user log in
# 此处不符合Restful风格，但其路由与admin get one user重合，本人解决JWT中间件筛选后无法解决其路由定位至admin get one user问题，故调整此处api
GET /auth/user/:id?password=""

# admin log in
GET /user/:id?password=""

# admin delete member
DELETE /user/:id

# admin get one user
GET /user/:id
 
# admin get all user
GET /user

# user update information
PUT /user/:id
{
	"name": str
	'password' str
	"email": str
	"phone": str
}
```



password在service层加密



### Util

`GenerateUUID`

`Encrypt`



-------------

`deleteUser`->`DeleteUserService` -> `DeleteMapper`

​																-> `FindMapper`

`deleteUser`根据true/false返回对应HTTP Status Code

`DeleteUserService`校验是否有删除权限。 若无，则返回false，否则返回true

`FindMapper`

`DeleteMapper`向数据库中删除数据





`getUserInfo` -> `GetUserInfoService` -> `FindMapper`

`GetUserInfoService` 检验是否有查看权限。若有，则返回true和具体数据，否则返回false

`FindMapper`



`updateUserInfo` -> `UpdateUserInfoService` -> `UpdateMapper`



------

email

`register` -> 



--------------

`ligon` -> `loginService` -> `FindMapper`

islog

`loginService` 返回bool判断成功与否