/**
*@Author: haoxiongxiao
*@Date: 2019/5/7
*@Description: CREATE GO FILE service
 */
package service

import (
	"time"
)

func HeartBeatTicker(mq *RabbitMq) {
	for {
		mq.Publish(HeartBeatTest())
		time.Sleep(60 * time.Second)
	}
}
