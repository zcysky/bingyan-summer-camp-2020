# 阶段一 API

#### 注册（普通用户）

```
POST /register-form?id= & code=
主体数据{"id":string , "pwd":string , "nickname":string ,"phone":string , "email":string}
在id为空时，需要邮箱发送验证码，返回生成的uid
id不为空，检查验证码，无误后向数据库添加用户
```

#### 登录

``` 
GET /token?id= &pwd=
```

#### 用户 更改信息

```
PUT /user/:id 
{"pwd":string , "nickname":string ,"phone":string , "email":string}
string为""不修改
不能修改id
```

#### 管理员 删除用户

``` 
DELETE /user/:id
id string
```

根据用户id删除用户

#### 管理员 获取用户信息

```
GET /user?all= &id=
all boolean
id string
```

all 为true ,不检查id参数，获取所有用户信息

all 为false ,检查id参数，获取id用户信息