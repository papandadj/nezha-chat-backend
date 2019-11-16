package handler

import (
	"github.com/papandadj/nezha-chat-backend/proto/friend"
	"github.com/papandadj/nezha-chat-backend/proto/user"
)

// PostResp .
type PostResp struct {
}

//PostSerializer .
func PostSerializer(resp *friend.PostResp) PostResp {
	return PostResp{}
}

// DeleteByUserIDResp .
type DeleteByUserIDResp struct {
}

//DeleteByUserIDSerializer .
func DeleteByUserIDSerializer(resp *friend.DelByUserIDResp) DeleteByUserIDResp {
	return DeleteByUserIDResp{}
}

//GetListResp .
type GetListResp struct {
	List []getListList `json:"list"`
}

type getListList struct {
	Img      string `json:"img"`
	Username string `json:"username"`
	ID       string `json:"id"`
}

//GetListSerializer .
func GetListSerializer(resp *user.GetListResp) GetListResp {
	wResp := GetListResp{}
	for _, item := range resp.List {
		wResp.List = append(wResp.List, getListList{
			Img:      item.Img,
			Username: item.Username,
			ID:       item.Id,
		})
	}

	return wResp
}
