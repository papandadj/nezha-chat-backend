package dao

import (
	"sync"

	"github.com/papandadj/nezha-chat-backend/pkg/gormplus"
	"github.com/papandadj/nezha-chat-backend/pkg/log"
	"github.com/papandadj/nezha-chat-backend/user-srv/conf"
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
	UserPost(username, password string) (user *ModelUser, err error)
	UserGetByUsername(username string) (user *ModelUser, ok bool, err error)
	UserLogin(username, password string) (user *ModelUser, ok bool, err error)
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
