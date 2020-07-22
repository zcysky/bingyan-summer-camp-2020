# API 文档

## 认证模式

Authorization 头 Bearer token 模式，即在请求头中包含以下字段：

```
Authorization: Bearer <access token>
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

API 前缀：`/api/v1`

### POST /event

添加一个事件。

请求：

如果不提醒，可以不包含 `remind_time` 和 `remind_interval`。

```json
{
    "desc": "事件描述",
    "user": "用户 QQ 号",
    "time": 123456789, // 事件发生时间
    "remind": true, // 是否提醒
    "remind_time": 123456789, // 开始提醒时间
    "remind_interval": 10 // 每次提醒间隔
}
```

响应：

```json
{
    "_id": "事件 ID"
}
```

### GET /event?_id=事件ID&user=用户QQ号&remind=1

根据条件查询事件。`remind=1` 时查询符合这些条件且需要提醒的事件。

请求：无

响应：

```json
{
    "result": [
        {
            "_id": "事件 ID",
            "desc": "事件描述",
            "user": "用户 QQ 号",
            "time": 123456789, // 事件发生时间
            "remind": true, // 是否提醒
            "remind_time": 123456789, // 开始提醒时间
            "remind_interval": 10 // 每次提醒间隔
        }
    ]
}
```

### DELETE /event?_id=事件ID

删除一个事件。

请求：无

响应：无

### PUT /event

修改一个事件。

请求：

请求中包含哪些字段就修改哪些，其中 `_id` 必须提供，且如果将 `remind` 修改为 `true`，则必须提供 `remind_time` 和 `remind_interval`。

```json
{
    "_id": "事件 ID",
    "desc": "事件描述",
    "user": "用户 QQ 号",
    "time": 123456789, // 事件发生时间
    "remind": true, // 是否提醒
    "remind_time": 123456789, // 开始提醒时间
    "remind_interval": 10 // 每次提醒间隔
}
```
