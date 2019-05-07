/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE service
 */
package service

import (
	"github.com/tidwall/gjson"
)

type Command struct {
	Recv chan []byte
	Send chan []byte
}

var ActionMap map[string]func(c *Command, data []byte) = map[string]func(c *Command, data []byte){
	"build":     Build,
	"heartbeat": HeartBeatRecv,
	"start":     StartOrStop,
}

func NewCommand() *Command {
	return &Command{Recv: make(chan []byte, 512), Send: make(chan []byte, 512)}
}

func (c *Command) Start() {
	data := <-c.Recv

	actionKey := gjson.GetBytes(data, "action").String()
	if callback, ok := ActionMap[actionKey]; ok {
		callback(c, data)
	}
}
