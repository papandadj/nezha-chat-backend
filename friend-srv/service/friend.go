package service

import (
	"context"
	"fmt"

	"github.com/papandadj/nezha-chat-backend/proto/friend"
)

//Post .
func (s *Service) Post(ctx context.Context, req *friend.PostReq, resp *friend.PostResp) (err error) {
	fmt.Println(req)
	return
}

//DelByUserID .
func (s *Service) DelByUserID(ctx context.Context, req *friend.DelByUserIDReq, resp *friend.DelByUserIDResp) (err error) {
	return
}

//CheckIsFriend .
func (s *Service) CheckIsFriend(ctx context.Context, req *friend.CheckIsFriendReq, resp *friend.CheckIsFriendResp) (err error) {
	return
}

//GetList .
func (s *Service) GetList(ctx context.Context, req *friend.GetListReq, resp *friend.GetListResp) (err error) {
	return
}
