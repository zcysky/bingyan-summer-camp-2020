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

#### 邮箱信息有误时

- Status: HTTP 400 Bad Request
- Body;

```json
{
    "result":	"Email address is invalid"
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

#### 密码错误时

- Status: HTTP 403 Forbidden
- Body:

```json
{
    "result":	"wrong password"
}
```

#### 用户不存在时

- Status: HTTP 403 Forbidden
- Body:

```json
{
    "result":	"wrong email address"
}
```

## 全部查询

### Request

- Method: **GET**
- URL: `/api/queryall/<limit><page>`
- Header: token

### Response

#### 成功查询

- Status: HTTP 200 OK
- Body:

```json
{
    "count":	1,
    "limit":	50,
    "page":		1,
    "users": [
        {
            "user_id":	"***",
            "username":	"***",
            "tel":		"***",
            "email":	"***"
        }
    ]
}
```

#### 非管理员身份

- Status: HTTP 403 Forbidden
- Body:

```json
{
    "result":	"Insufficient permission"
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
    "user_id":	"***",
    "username":	"***",
    "tel":		"***",
    "email":	"***"
}
```



#### 非管理员身份

- Status: HTTP 403 Forbidden
- Body:

```json
{
    "result":	"Insufficient permission"
}
```

#### 查找不到

- Status: HTTP 404 Not Found
- Body:

```json
{
    "result":	"Not found the user"
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
    "result":	"Delete the user successfully",
    "user_id":	"***",
    "username":	"***",
    "tel":		"***",
    "email":	"***"
}
```



#### 非管理员身份

- Status: HTTP 403 Forbidden
- Body:

```json
{
    "result":	"Insufficient permission"
}
```

#### 查找不到

- Status: HTTP 404 Not Found
- Body:

```json
{
    "result":	"Not found the user"
}
```



