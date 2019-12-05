package service

import (
	"github.com/papandadj/nezha-chat-backend/common-srv/dao"
	"github.com/papandadj/nezha-chat-backend/pkg/log"
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
