/**
*@Author: haoxiongxiao
*@Date: 2019/5/7
*@Description: CREATE GO FILE service
 */
package service

import (
	"github.com/spf13/viper"
	"time"
)

func HeartBeatTicker(mq *RabbitMq) {
	for {
		mq.Publish(HeartBeatTest())
		time.Sleep(time.Duration(viper.GetDuration("heartbeat.ticker_time") * time.Second))
	}
}
