# 数据库结构文档

## Mongo

数据库名：`qq-bot-ref`

### 事件集合 event

- _id `ObjectId`：事件 ID
- user `uint`：事件提醒的 QQ 号
- desc `String`：事件描述
- time `int64`：事件发生时间（时间戳）
- remind `Boolean`：事件是否提醒
- remind_time `int64`：事件提醒开始时间（时间戳）
- remind_interval `int`：事件提醒间隔（分钟）
