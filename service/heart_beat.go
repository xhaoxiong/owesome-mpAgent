/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE service
 */
package service

import (
	"encoding/json"
	"github.com/spf13/viper"
	"mpAgent/models"
)

func HeartBeat() []byte {
	var hresult models.HeartBeatResult

	hresult.Action = "heartbeat"
	hresult.Data.Ip = viper.GetString("heart_beat.ip")
	hresult.Data.Queue = viper.GetString("heart_beat.queue")
	bytes, _ := json.Marshal(hresult)
	return bytes
}
