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
	Dao dao.Interface
}

//New .
func New(daoIns dao.Interface) (service *Service) {
	service = new(Service)
	service.Dao = daoIns
	return
}
