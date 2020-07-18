```
db.user.find({"_id":ObjectId("5f0eb47dd28ecc4fdbe1fdc7")}).pretty()
```

```
db.user.update({"_id":ObjectId("5f0eb47dd28ecc4fdbe1fdc7")},{$set:{"password":NumberInt(123)}})
```

```
db.user.deleteMany({username:"yokel3"})
```

```
db.user.find({username:/y\w*k\w*l1/})   匹配yokel1  y*k*l1
```

```
db.getCollection('user').find({$or:[{friends:"text1"},{friends:"text2"}]}).count()
```

```
db.user.update({},[{"$set":{"class_name":{"$concat":["$school"," ","$major"," ","$class"]}}}],{multi:true})
```

