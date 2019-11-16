package dao

import (
	"sync"

	"github.com/papandadj/nezha-chat-backend/friend-srv/conf"
	"github.com/papandadj/nezha-chat-backend/pkg/gormplus"
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
	Post(tokenID, userID string) (friend *ModelFriend, err error)
	DeleteByUserID(tokenID, userID string) (err error)
	GetList(tokenID string) (friend []*ModelFriend, err error)
	checkIsFriend(tokenID, userID string) (result bool, err error)
}

//Dao .
type Dao struct {
	DB *gormplus.DB
}

//Ping .
func (a *Dao) Ping() (err error) {
	err = a.DB.DB.DB().Ping()
	return
}

//Init init a dao
func Init() {
	lock.Lock()
	defer lock.Unlock()

	if dao != nil {
		logger.Infoln("Dao 已经初始化过了")
		return
	}

	dao = &Dao{
		DB: newDB(),
	}

	err := dao.Ping()
	if err != nil {
		logger.Fatalln(err)
	}
	return
}

//NewDB 返回gorm链接实例
func newDB() *gormplus.DB {
	c := conf.GetGlobalConfig()
	return gormplus.New(&gormplus.Config{
		Debug: c.MySQL.Debug,
		DSN:   c.MySQL.DSN(),
	})
}

//GetDao .
func GetDao() *Dao {
	return dao
}
