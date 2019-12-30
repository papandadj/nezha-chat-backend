package handler

import (
	"github.com/papandadj/nezha-chat-backend/proto/friend"
)

// PostResp .
type PostResp struct {
}

//PostSerializer .
func PostSerializer(resp *friend.PostResp) PostResp {
	return PostResp{}
}
