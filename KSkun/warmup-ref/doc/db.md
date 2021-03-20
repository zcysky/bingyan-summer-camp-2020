# 数据库结构文档

## Mongo

### 用户集合 user

- _id `ObjectId`：用户 ID
- username `String`: 用户名
- password `String`：密码
- phone `String`：手机号
- email `String`：邮箱
- is_admin `Boolean`：是否有管理权限
- verified `Boolean`：是否已经验证邮箱

## Redis

- `code:<邮箱验证码>`：验证码对应的用户 ID

