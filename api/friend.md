# 用户接口文档

注： 未经请求全部为JSON格式， 方式为POST

```
baseUrl: http://localhost:9501
```

## 目录:

[1. 用户注册](#sign_up)
[2. 用户登录](#login)
[3. 根据查询获取用户列表](#get_list)


## 接口列表:

### sign_up

#### 请求URL:

```
http://
```

#### 请求参数

|参数|是否必选|类型|说明|
|:-----|:-------:|:-----|:-----|
|username      |Y       |string  |用户名称
|password      |Y       |string  |用户密码


成功返回:

```json
{

}
```

失败返回: 

```json
{
    "Msg":"用户已经注册",
    "Code": 409
}
```


### login

#### 请求URL:

```
http://
```

#### 请求参数

|参数|是否必选|类型|说明|
|:-----|:-------:|:-----|:-----|
|username      |Y       |string  |用户名称
|password      |Y       |string  |用户密码


成功返回:

status: 200

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJ0aW1lc3RhbXAiOjEyMzQ1LCJ1c2VybmFtZSI6Im5lemhhIn0.xHGOIRzIylTLWx-ceTa6UMsw4uO-kQk4asfZoT0XKms"
}
```

失败返回: 

```json
{
    "Msg": "用户账号或者密码错误",
    "Code": 401
}
```


### get_list

#### 请求URL:

```
http://
```

#### 请求参数

|参数|是否必选|类型|说明|
|:-----|:-------:|:-----|:-----|
|name      |N      |string  |根据名称模糊搜索
|ids      |Y       |array  |需要查找的用户id


成功返回:

```json
{
    "list": [
        {
            "img": "http://....",
            "username": "nezhe",
            "id": "1"
        }
    ]
}
```


