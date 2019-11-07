package service

import (
	"github.com/papandadj/nezha-chat-backend/auth-srv/conf"
	"github.com/papandadj/nezha-chat-backend/auth-srv/dao"
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
	Dao           dao.Interface
	defaultExpire int
	tokenSecrete  string
}

//New .
func New(daoIns dao.Interface) (service *Service) {
	cfg := conf.GetGlobalConfig()
	service = new(Service)
	service.Dao = daoIns
	service.defaultExpire = cfg.Memory.DefaultExpiration
	service.tokenSecrete = cfg.Secrete
	return
}
