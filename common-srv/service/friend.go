package service

import (
	"context"
	"fmt"

	"github.com/papandadj/nezha-chat-backend/common-srv/conf"

	"github.com/papandadj/nezha-chat-backend/proto/common"
)

//GetList .
func (s *Service) GetList(ctx context.Context, req *common.GetListReq, resp *common.GetListResp) (err error) {
	userImages, err := s.Dao.UserImgGetList()
	if err != nil {
		return
	}

	resp.List = make([]string, 0)
	for _, userImage := range userImages {
		resp.List = append(resp.List, fmt.Sprintf("%s%s", conf.GetGlobalConfig().ImgPrefix, userImage.URL))
	}

	return
}
