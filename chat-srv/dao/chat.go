package dao

import (
	"github.com/papandadj/nezha-chat-backend/chat-srv/conf"
	"github.com/streadway/amqp"
)

// Post .
func (d *Dao) Post(queue string, data []byte) (err error) {
	cfg := conf.GetGlobalConfig()
	err = d.RabbitChannel.Publish(cfg.RabbitMq.ChatExchangeName, queue, true, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
	})
	return
}

//CreateQueue .
func (d *Dao) CreateQueue(queue string) (err error) {
	cfg := conf.GetGlobalConfig()
	_, err = d.RabbitChannel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return
	}

	err = d.RabbitChannel.QueueBind(queue, queue, cfg.RabbitMq.ChatExchangeName, false, nil)
	return
}
