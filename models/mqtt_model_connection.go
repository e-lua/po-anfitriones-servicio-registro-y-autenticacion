package models

import (
	"github.com/streadway/amqp"
)

type MqttDB struct {
	amqp.Config
}

var MqttCN = GetConnMqtt()

func GetConnMqtt() *amqp.Connection {

	conn, error_connec_mqtt := amqp.Dial("amqp://edwardlopez:servermqtt@143.110.233.233:8888/")

	if error_connec_mqtt != nil {

		return conn
	}

	return conn
}
