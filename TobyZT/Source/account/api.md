# API

#### 登陆

- POST /api/v1/login 登陆

- 接受JSON数据，返回JSON数据

  接受JSON数据示例：

  ```json
  {
      "email": ""
      "password": "19260817",
  }
  ```
  
  返回JSON数据示例：

  ```json
  {
	    "message": "success",
      "status": 200,
      "userid": "5f131a71252025084b9c1cc7"
      "admin": 1,
      "Authorization": "bearer xxx"
  } 
  ```

  >  status 200 表示成功，401表示账号或密码错误，410表示用户不存在



#### 注册

- POST /api/v1/signup 注册

- 接受JSON数据，返回JSON数据

  接受JSON数据示例：
  
  ```json
  {
      "username": "yyhtql",
      "password": "",
      "phone": "",
      "email": "",
      "token": ""
  }
  ```
  
  > token为空时，后端会向邮箱发送验证邮件，验证邮件包含token
  >
  > token不为空时，后端将对token进行验证，将数据存入数据库
  
  返回JSON数据示例：
  
  ```json
  {
  	"message": "success",
  	"status": 201,
  	"Authorization": "bearer xxx"
  }
  ```
  
  >  status 200表示发送验证邮件成功，201 表示注册成功，400表示用户填写信息有误，401表示token错误，403表示用户已存在



#### 管理员获取所有成员信息

- GET /api/v1/users

- 接受JSON数据，请求数据示例：

  ```json
  {
      "limit": 60,
      "page": 3
  }
  ```

  >  `limit`表示单条JSON数据返回的最大成员数，`page`表示查询的页码，

- 返回JSON数据，返回数据示例：

  ```json
  {
      "message": "success",
      "status": 200,
      "total": 114,
      "page": 6,
      "limit": 60,
      "users": [
       {
      	"userid": 123,
      	"username": "user1",
      	"phone": "",
      	"email": ""
  	 },
  	 {
          "userid": 456,
      	"username": "user2",
      	"phone": "",
      	"email": ""
       }
      ]
  }
  ```

  >  status 200表示成功，404表示该页数据不存在，401表示认证失败
  >
  >  `total`表示查询到的成员总数，`page`表示当前页码，
  >
  >  `limit`表示单条JSON数据返回的最大成员数

- 请求时需在请求头加入Authorization，否则会认证失败，引发status 401



#### 管理员获取某个成员信息

- GET /api/v1/users/<user_id>

- `user_id`表示要查询的成员的ID

- 返回JSON数据，返回数据示例：

  ```json
  {
  	"message": "success",
      "status": 200,
      "userinfo": {
          "userid": 114514,
          "username": "yyhtql",
          "phone": "",
          "email": ""
      }
  }
  ```

  >  status 200表示成功，404表示查询不到该成员，401表示认证失败

- 请求时需在请求头加入Authorization，否则会认证失败，引发status 401



#### 管理员删除普通用户

- DELETE /api/v1/users/<user_id>

- `user_id`表示要删除成员的ID

- 返回JSON数据，返回数据示例：

  ```json
  {
  	"message": "success",
      "status": 204
  }
  ```

  >  status 204表示删除成功，401表示认证失败，403表示删除失败

- 请求时需在请求头加入Authorization，否则会认证失败，引发status 401



#### 用户修改信息

- PUT /api/v1/users

- 接受JSON数据，返回JSON数据

  请求JSON数据示例：

  ```json
  {
      "username": "yyhtql",
      "password": "",
      "phone": "",
      "email": "",
  }
  ```

  >  请求时需在请求头加入Authorization，否则会认证失败，引发status 401

  	返回JSON数据示例：

  ```json
  {
      "message": "success",
      "status": 201
  }
  ```

  > status 201表示修改成功，401表示认证失败
  >
  > 修改成功后，JSON会附有新的JWT令牌


