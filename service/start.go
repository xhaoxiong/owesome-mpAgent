/**
*@Package:service
*@Author: haoxiongxiao
*@Date: 2019/5/7
*@Description: create go file in service package
 */
package service

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

func StartOrStop(c *Command, data []byte) {
	params := gjson.GetBytes(data, "params")

	port := params.Get("port").String()
	serviceName := params.Get("service_name").String()
	options := optionsReplace(params.Get("options").String(), port)
	version := params.Get("version").String()
	cmds := gjson.GetBytes(data, "cmds").Array()
	taskId := gjson.GetBytes(data, "taskId").String()
	fmt.Println("开始遍历cmds...")
	for index := range cmds {
		cmdStr := cmds[index].Get("cmd").String()
		dir := cmds[index].Get("dir").String()

		cmdStr = cmdReplace(cmdStr, "", serviceName, "", version, port, options, "")
		splits := strings.Split(cmdStr, " ")
		sendResult(c, index, splits, dir, taskId, cmdStr)
	}
	fmt.Println("遍历cmds完成...")
	c.Send <- ReturnEnd(taskId)

}
