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

func HeartBeatTest() []byte {
	var hresult models.HeartBeatResult

	hresult.Action = "heartbeat"
	hresult.Data.Ip = viper.GetString("heart_beat.ip")
	hresult.Data.Queue = viper.GetString("heart_beat.queue")
	bytes, _ := json.Marshal(hresult)

	return bytes
}

func HeartBeat(c *Command, data []byte) {
	var heartbeat models.HeartBeatResult

	heartbeat.Action = "heartbeat"
	heartbeat.Data.Queue = viper.GetString("heart_beat.queue")
	heartbeat.Data.Ip = viper.GetString("heart_beat.ip")
	heartbeat.Data.User = viper.GetString("heart_beat.user")
	heartbeat.Data.Passwd = viper.GetString("heart_beat.password")
	heartbeat.Data.FreeCpu = FreeCPU()
	heartbeat.Data.FreeDisk = FreeDisk()
	heartbeat.Data.FreeMem = FreeRAM()
	heartbeat.Data.DiskInfo = CPUCheck()
	heartbeat.Data.MemInfo = RAMCheck()
	heartbeat.Data.DiskInfo = DiskCheck()
	heartbeat.Data.MachineType = viper.GetStringSlice("heart_beat.machine_type")
	bytes, _ := json.Marshal(heartbeat)

	c.Send <- bytes
}
