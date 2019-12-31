# 朋友接口

注： 未经请求全部为JSON格式， 方式为POST

```
baseUrl: http://localhost:9501
```

## 目录:

[1. 添加朋友](#post)
[2. 删除朋友](#delete/:user_id)
[3. 获取朋友](#get_list)



## 接口列表:

### post 

#### 请求URL:

```
http://
```

#### 请求参数 (token)

|参数|是否必选|类型|说明|
|:-----|:-------:|:-----|:-----|
|user_id      |Y       |string  |用户id


成功返回:

```json
{

}
```

### delete/:user_id

#### 请求URL:

```
http://
```

#### 请求参数 (token)

|参数|是否必选|类型|说明|
|:-----|:-------:|:-----|:-----|
|user_id(param)      |Y       |string  |用户id


成功返回:


```json
{
}
```

### get_list

#### 请求URL:

```
http://
```

#### 请求参数 (token)

|参数|是否必选|类型|说明|
|:-----|:-------:|:-----|:-----|


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


