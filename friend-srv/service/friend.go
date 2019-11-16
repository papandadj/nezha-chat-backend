package service

import (
	"context"

	"github.com/papandadj/nezha-chat-backend/proto/friend"
)

//Post 小的id在前面， 大的在后面.
func (s *Service) Post(ctx context.Context, req *friend.PostReq, resp *friend.PostResp) (err error) {
	user1, user2 := compareTwoUserPosition(req.TokenId, req.UserId)

	_, err = s.Dao.Post(user1, user2)
	if err != nil {
		logger.Errorln(err)
		return
	}

	return
}

//DelByUserID .
func (s *Service) DelByUserID(ctx context.Context, req *friend.DelByUserIDReq, resp *friend.DelByUserIDResp) (err error) {
	user1, user2 := compareTwoUserPosition(req.TokenId, req.UserId)

	err = s.Dao.DeleteByUserID(user1, user2)
	return
}

//CheckIsFriend .
func (s *Service) CheckIsFriend(ctx context.Context, req *friend.CheckIsFriendReq, resp *friend.CheckIsFriendResp) (err error) {
	return
}

//GetList .
func (s *Service) GetList(ctx context.Context, req *friend.GetListReq, resp *friend.GetListResp) (err error) {
	friends, err := s.Dao.GetList(req.TokenId)
	if err != nil {
		return
	}

	for _, friend := range friends {
		if friend.UserID1 == req.TokenId {
			resp.List = append(resp.List, friend.UserID2)
		} else {
			resp.List = append(resp.List, friend.UserID1)
		}

	}

	return
}

//compareTwoUserPosition .
func compareTwoUserPosition(id1, id2 string) (user1, user2 string) {
	if id1 > id2 {
		user1 = id2
		user2 = id1
	} else {
		user1 = id1
		user2 = id2
	}
	return
}
