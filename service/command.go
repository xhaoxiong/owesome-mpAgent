/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE service
 */
package service

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
)

type Command struct {
	Recv chan []byte
	Send chan []byte
}

var ActionMap map[string]func(c *Command, data []byte) = map[string]func(c *Command, data []byte){
	"build": Build,
}

func NewCommand() *Command {
	return &Command{Recv: make(chan []byte, 512), Send: make(chan []byte, 512)}
}

func (c *Command) Start() {
	data := <-c.Recv

	var t interface{}

	gjson.Unmarshal(data, &t)
	data = []byte(cast.ToString(t))

	actionKey := gjson.GetBytes(data, "action").String()
	cmds := gjson.GetBytes(data, "cmds").Array()

	fmt.Println("cmds:", cmds)
	fmt.Println("key:", actionKey)
	if callback, ok := ActionMap[actionKey]; ok {
		callback(c, data)
	}
}
