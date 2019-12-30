package dao

import (
	"sync"

	"github.com/streadway/amqp"

	"github.com/papandadj/nezha-chat-backend/chat-srv/conf"

	"github.com/papandadj/amqp/connector"

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
	Post(queue string, data []byte) (err error)
	CreateQueue(queue string) (err error)
}

//Dao .
type Dao struct {
	RabbitConn    *connector.Connection
	RabbitChannel *connector.Channel
}

//Ping .
func (a *Dao) Ping() (err error) {
	return
}

//Init init a dao
func Init() {
	lock.Lock()
	defer lock.Unlock()

	cfg := conf.GetGlobalConfig()
	if dao != nil {
		logger.Infoln("Dao 已经初始化过了")
		return
	}

	rabbitConn, err := newRabbitConn()
	if err != nil {
		panic(err)
	}
	rabbitChannel, err := rabbitConn.Channel()
	if err != nil {
		panic(err)
	}
	logger.With("url", cfg.RabbitMq.DSN()).Infoln("RabbitMq 链接初始化成功")
	dao = &Dao{
		RabbitChannel: rabbitChannel,
		RabbitConn:    rabbitConn,
	}

	//初始化聊天队列
	err = rabbitChannel.ExchangeDeclare(cfg.RabbitMq.ChatExchangeName, amqp.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		logger.Fatalln(err)
	}

	err = dao.Ping()
	if err != nil {
		logger.Fatalln(err)
	}
	return
}

//GetDao .
func GetDao() *Dao {
	return dao
}

//NewRabbitConn 初始化rabbit链接
func newRabbitConn() (*connector.Connection, error) {
	c := conf.GetGlobalConfig()
	return connector.Dial(c.RabbitMq.DSN())
}
