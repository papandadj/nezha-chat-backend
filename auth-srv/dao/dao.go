package dao

import (
	"sync"
	"time"

	"github.com/papandadj/nezha-chat-backend/auth-srv/conf"

	"github.com/patrickmn/go-cache"

	"github.com/papandadj/nezha-chat-backend/pkg/log"
)

var (
	logger log.Logger
	dao    *Dao
	lock   sync.Mutex
)

func init() {
	logger = log.Base()
}

//Interface dao
type Interface interface {
	AuthCheck(token string) (success bool, err error)
	AuthSaveToken(username, token string) (err error)
	AuthDelToken(token string) (err error)
}

//Dao .
type Dao struct {
	Mem *cache.Cache
}

//Init init a dao
func Init() {
	cfg := conf.GetGlobalConfig()
	lock.Lock()
	defer lock.Unlock()

	if dao != nil {
		logger.Infoln("Dao 已经初始化过了")
		return
	}
	expiration := time.Duration(cfg.Memory.DefaultExpiration)
	intervalClear := time.Duration(cfg.Memory.IntervalClear)

	dao = &Dao{
		Mem: cache.New(time.Minute*expiration, time.Minute*intervalClear),
	}

	return
}

//GetDao .
func GetDao() *Dao {
	return dao
}
