有的接口需要权限才能， 使用jwt进行认证， 需要在http的header添加`token`字段。

如下面的`postman`请求, 接口请求参数后面加`(token)`表示该接口需要`token`

```txt
POST /get_list HTTP/1.1
Host: localhost:9500
Content-Type: application/json
token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJ0aW1lc3RhbXAiOjE1NzM3ODgyMjEsInVzZXJuYW1lIjoibmV6aGEifQ.p25qAekgagMuBl-BCwNryJtNzQln694W_wO24PdhzTs
Cache-Control: no-cache
Postman-Token: fc6775bb-d1e0-5a77-8a6f-12b55f5d89c0

{
	"name":"ne"
}
```