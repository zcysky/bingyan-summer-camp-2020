添加事件

```
/save 保存当前事件，在消息中包含该命令
```

查询所有事件

```
/showMyEvent 显示已保存的所有事件，在消息中包含该命令
```

删除事件

```
/delete <id>      删除eventid为id的事件，该命令不能包含其它字符
```



添加提醒

```
/addNoti <id> <year>-<month>-<day>-<hour>-<minute> <interval> <advance>
将id事件加入提醒库中，collection为qqid+"Noti"

year,month,day为0表示当时年月日
{
	"eventId":
	"time":unix timestamp
	"interval"//min
	"advance"//min
}
```

