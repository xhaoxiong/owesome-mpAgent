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
	"github.com/xhaoxiong/util"
	"os"
	"strings"
)

func Build(c *Command, data []byte) {

	params := gjson.GetBytes(data, "params")
	repoUrl := params.Get("repo_url").String()
	serviceName := params.Get("service_name").String()
	Profile := params.Get("profile").String()
	Version := params.Get("version").String()
	Port := params.Get("port").String()
	Options := optionsReplace(params.Get("options").String(), Port)
	workDir := params.Get("work_dir").String()
	cmds := gjson.GetBytes(data, "cmds").Array()
	taskId := gjson.GetBytes(data, "taskId").String()

	if !util.Exists(workDir) {
		os.MkdirAll(workDir, 0755)
	}

	fmt.Println("开始遍历cmds...")
	for index := range cmds {
		dir := cmds[index].Get("dir").String()
		cmdStr := cmds[index].Get("cmd").String()

		dir = strings.Replace(dir, "$work_dir", workDir, -1)
		dir = strings.Replace(dir, "$service_name", serviceName, -1)
		if !util.Exists(dir) {
			os.MkdirAll(dir, 0755)
		}
		cmdStr = cmdReplace(cmdStr, repoUrl, serviceName, Profile, Version, Port, Options, workDir)
		splits := strings.Split(cmdStr, " ")
		sendResult(c, index, splits, dir, taskId, cmdStr)
	}
	fmt.Println("遍历cmds完成...")
	c.Send <- ReturnEnd(taskId)

}

func sendResult(c *Command, index int, splits []string, dir, taskId, cmdStr string) {
	if result, err := Excute(splits[0], splits[1:], dir); err != nil {
		returnResult := ReturnResult(index, taskId, -1, cast.ToString(err), cmdStr)

		c.Send <- returnResult

	} else {
		returnResult := ReturnResult(index, taskId, 0, string(result), cmdStr)
		c.Send <- returnResult
	}
}
