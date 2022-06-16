package rabbitmq

import (
	"fmt"
	"log"

	"github.com/OswinZheng/gin-web-F/configs"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel

func Conn() *amqp.Connection {
	if conn == nil {
		var err error
		log.Printf(fmt.Sprintf("amqp://%s:%s@%s:%d/", configs.Get().Rabbitmq.User, configs.Get().Rabbitmq.Password, configs.Get().Rabbitmq.Host, configs.Get().Rabbitmq.Port))
		conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", configs.Get().Rabbitmq.User, configs.Get().Rabbitmq.Password, configs.Get().Rabbitmq.Host, configs.Get().Rabbitmq.Port))
		if err != nil {
			panic(err)
		}
	}
	return conn
}

func Channel() *amqp.Channel {
	if channel == nil {
		var err error
		channel, err = Conn().Channel()
		if err != nil {
			panic(err)
		}
	}
	return channel
}
