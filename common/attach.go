package common

//AttachAuth 使用中间件解析完成后， token的结构体
type AttachAuth struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

//AttachAuthKey 存入到context 的key值
const AttachAuthKey = "auth"
