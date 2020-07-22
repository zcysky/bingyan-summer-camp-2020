# API 文档

## 认证模式

Authorization 头 Bearer 令牌模式，即在请求头中包含以下字段：

```
Authorization: Bearer <JWT token>
```

## 数据交换模式

在所有的 POST 请求中使用 JSON 格式作为请求体格式，所有服务端响应使用 JSON 格式。响应格式如下：

```json
{
    "success": true,
    "hint": "",
    "data": {} // 根据接口实际情况决定
}
```

下列接口定义中，服务器响应格式均指 data 中的子文档格式。

如果遇到错误，则 `success` 值一定为 `false`，且 `hint` 字段中包含错误的具体内容。

## API 定义

带 \* 号标注的接口为需要认证的接口，带 \*\* 号标注的为需要管理权限的接口。

API 前缀：`/api/v1`

### GET /user/token?username=\<用户名\>&password=\<密码\>

获取用户 JWT 令牌。

请求：无

响应：

```json
{
    "token": "令牌",
    "expire_time": 123456789 // 令牌到期时间
}
```

### POST /user/info

新用户注册，全部字段都不可为空。

请求：

```json
{
    "username": "用户名",
    "password": "密码",
    "phone": "手机号",
    "email": "邮箱地址"
}
```

响应：

```json
{
    "_id": "用户 ID"
}
```

### POST /user/verify

确认邮箱。

*注：本接口定义不符合 RESTful 规范。*

请求：

```json
{
    "_id": "用户 ID",
    "code": "邮箱验证码"
}
```

响应：无

### \*PUT /user/info

修改用户个人信息。可以修改的字段有：用户名（不可重复）、密码、手机号、邮箱地址（需要重新确认），全部字段都不可为空。请求体中包含哪些字段就修改哪些。

请求：

```json
{
    "_id": "用户 ID",
    "username": "用户名",
    "password": "密码",
    "phone": "手机号",
    "email": "邮箱地址"
}
```

响应：无

### \*\*DELETE /user?_id=\<用户 ID\>

删除用户。

请求：无

响应：无

### \*GET /user/info?_id=\<用户 ID\>

获取用户信息，不包含密码，用户 ID 可以为空以获取所有人的信息。只有管理权限能获取其他人的用户信息或所有人的信息。

请求：无

响应：

```json
{
    "result": [
        {
            "_id": "用户 ID",
            "username": "用户名",
            "phone": "手机号",
            "email": "邮箱地址"
        }
        // ... 以下略
    ]
}
```

