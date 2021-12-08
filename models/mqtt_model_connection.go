package models

import (
	"sync"

	"github.com/streadway/amqp"
)

type MqttDB struct {
	amqp.Config
}

var MqttCN = GetConnMqtt()

var (
	once_mqtt sync.Once
	p_mqtt    *amqp.Connection
)

func GetConnMqtt() *amqp.Connection {

	once_mqtt.Do(func() {
		p_mqtt, _ = amqp.Dial("amqp://edwardlopez:servermqtt@143.110.233.233:8888/")

	})

	return p_mqtt
}
