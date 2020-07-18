```javascript
//插入
db.user.insertMany([
{user_name: "Bob", password:"bob", reg_time: new Date(), email:"bbb@bbb.com"},
{user_name: "Saber", password:"saber", reg_time: new Date(), email:"type@moon.com"},
{user_name: "Lancer", password:"lancer", reg_time: new Date(), email:"type@moon.com"},
{user_name: "Arch", password:"arch", reg_time: new Date(), email:"type@moon.com"}
]);
```

```javascript
// 查找
db.user.find({_id: ObjectId("5f0eb1fbfe21f29bdba0b8c6")})

// 重新填充
db.user.replaceOne(
{_id: ObjectId("5f0eb1fbfe21f29bdba0b8c6")},
{user_name: "Caster"}
)

// 修改
db.user.updateOne(
{_id: ObjectId("5f0eb45bfe21f29bdba0b8c8")},
{$set:{user_name: "Saber Lily"}}
)

//删除
db.user.deleteOne({_id : ObjectId("5f0eb1fbfe21f29bdba0b8c6")})

//模糊查找 $regex：正则匹配
db.user.find({ user_name : { $regex: "a"} })
```





```javascript
db.user.find(
    { 
        $or : [
             { friends:{$regex:"test1"} },
             { friends:{$regex:"test2"} } 
        ]
    }
).count()
```

```javascript
db.getCollection("user").find({}).forEach(
	function(item){
        db.getCollection("user").update(
            {_id: item._id}, 
            {$set : {class_name : item.school + item.major + item.class}}
        )
    }
)
```

