# 用户管理系统

## 注册

### Request

- Method: **POST**
- URL: `/api/register/`
- Body:

```json
{
    "is_admin":		0/1,
    "invitation":	"***",
    "username":		"***",
    "password":		"***",
    "tel":			"***",
    "email":		"***"
}
```

`is_admin`为0则注册为普通用户，为1则检验邀请码`invitation`，邀请码正确则注册为管理员。

### Response

#### 注册成功时

- Status: HTTP 200 OK
- Body:

```json
{
    "result":	"Registered successfully",
    "id":		"***"
}
```

`id`为用户信息在MongoDB中存储的id

#### 邮箱已被注册时

- Status: HTTP 403 Forbidden
- Body:

```json
{
    "result":	"This Email address has been registered"
}
```

#### 邮箱信息有误时

- Status: HTTP 400 Bad Request
- Body;

```json
{
    "result":	"Email address is invalid",
    "error":	"报错内容"
}
```

#### 管理员注册邀请码不正确时

- Status: HTTP 400 Bad Request
- Body;

```json
{
    "result":	"invitation code wrong"
}
```

## 登录

### Request

- Method: **POST**
- URL:`/api/login/`
- Body:

```json
{
    "is_admim":		0/1,
    "email":		"***",
    "password":		"***"
}
```

`is_admin`为0则从普通用户表中查找信息登录，为1则从管理员表中操作。

### Response

#### 登录成功时

- Status: HTTP 200 OK
- Body:

```json
{
	"result":	"Login successfully",
	"Authorization": "Bearer " + token
}
```

#### 密码错误或用户不存在时

- Status: HTTP 400 Bad Request
- Body:

```json
{
    "result":	"Wrong password or wrong email address",
    "error":	"报错内容"
}
```

#### 登录信息格式错误时

- Status: HTTP 400 Bad Request
- Body:

```json
{
	"result":	"Wrong form struct",
    "error":	"报错内容"
}
```



## 全部查询

### Request

- Method: **POST**
- URL: `/api/queryall/`
- Header: token
- Body:

```json
{
    "limit": 	"每页显示的用户数量",
    "page":		"第几页"
}
```

### Response

#### 成功查询

- Status: HTTP 200 OK
- Body:

```json
{
    "result":	"Query all users successfully",
    "total":	"普通用户总人数",
    "limit":	50,
    "page":		1,
    "users": [
        {
            "user_id":	"***",
            "user_name":	"***",
            "tel":		"***",
            "email":	"***"
        }
    ]
}
```

#### 查询失败

- Status: HTTP 403Forbidden 或 400 Bad Request
- Body:

```json
{
    "result":	"失败结果",
    "error":	"具体报错"
}
```

## 单人查询

### Request

- Method: **GET**
- URL: `/api/query/<user_id>`
- Header: token

### Response

#### 成功查询

- Status: HTTP 200 OK
- Body:

```json
{
    "result":	"query successfully",
    "user":	{
        "user_id":		user.UserID,
        "user_name":	user.Username,
        "user_phone":	user.Phone,
        "user_email":	user.Email,
	}
}
```

#### 查询失败

- Status: HTTP 403 Forbidden 或 400 Bad Request
- Body:

```json
{
    "result":	"失败结果",
    "error":	"具体报错"
}
```

## 删除

### Request

- Method: **GET**
- URL: `/api/delete/<user_id>`
- Header: token

### Response

#### 成功删除

- Status: HTTP 200 OK
- Body:

```json
{
    "result":	"remove successfully",
    "user_id":	"***",
}
```

#### 删除失败

- Status: HTTP 403 Forbidden 或 400 Bad Request
- Body:

```json
{
    "result":	"失败结果",
    "error":	"具体报错"
}
```



