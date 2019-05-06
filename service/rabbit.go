/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE service
 */
package service

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"time"
)

var err error

type RabbitMq struct {
	rabbitConn *amqp.Connection
}

func NewRabbitMq() *RabbitMq {
	return &RabbitMq{rabbitConn: connect()}
}

func connect() (conn *amqp.Connection) {
	username := viper.GetString("rabbitmq.username")
	password := viper.GetString("rabbitmq.password")
	addr := viper.GetString("rabbitmq.addr")

	conn, err = amqp.Dial("amqp://" + username + ":" + password + "@" + addr + "/")
	if err != nil {
		panic(err)
	}
	return
}

func (this *RabbitMq) Publish(info []byte) {
	c, err := this.rabbitConn.Channel()
	if err != nil {
		panic(err)
	}
	if err = c.ExchangeDeclare(viper.GetString("rabbitmq.producer.exchange"), viper.GetString("rabbitmq.producer.kind"),
		true, false, false, false, nil); err != nil {
		panic(err)
	}

	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/json",
		Body:         info,
	}
	if err = c.QueueBind(viper.GetString("rabbitmq.producer.queue"), viper.GetString("rabbitmq.producer.queue"),
		viper.GetString("rabbitmq.producer.exchange"), false, nil); err != nil {
		panic(err)
	}

	err = c.Publish(viper.GetString("rabbitmq.producer.exchange"),
		viper.GetString("rabbitmq.producer.queue"), false,
		false, msg)
	if err != nil {
		panic(err)
	}
	c.Close()

}

func (this *RabbitMq) Consumer() (s []byte) {

	c, err := this.rabbitConn.Channel()
	if err != nil {
		panic(err)
	}
	defer c.Close()
	msg, ok, err := c.Get(viper.GetString("rabbitmq.consumer.queue"), false)
	if err != nil {
		panic(err)
	}

	if !ok {
		return nil
	}

	err = c.Ack(msg.DeliveryTag, false)

	s = msg.Body
	return
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
