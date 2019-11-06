package service

import (
	"github.com/papandadj/nezha-chat-backend/pkg/log"
	"github.com/papandadj/nezha-chat-backend/user-srv/dao"
)

var (
	logger log.Logger
)

func init() {
	logger = log.Base()
}

//Service .
type Service struct {
	Dao *dao.Dao
}

//New .
func New() (service *Service) {
	dao.Init()
	service = new(Service)
	service.Dao = dao.GetDao()

	return
}
