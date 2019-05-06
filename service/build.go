/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE service
 */
package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"github.com/xhaoxiong/util"
	"mpAgent/models"
	"os"
	"os/exec"
	"strings"
)

func Build(c *Command, data []byte) {

	params := gjson.GetBytes(data, "params")
	repoUrl := params.Get("repo_url").String()
	serviceName := params.Get("service_name").String()
	Profile := params.Get("profile").String()
	Version := params.Get("version").String()
	Port := params.Get("port").String()
	Options := params.Get("options").String()
	workDir := params.Get("work_dir").String()

	Options = strings.Replace(Options, "$HOST", os.Getenv("HOST"), -1)
	Options = strings.Replace(Options, "$REGISTER_CENTER", os.Getenv("REGISTER_CENTER"), -1)

	if !util.Exists(workDir) {
		os.MkdirAll(workDir, 0755)
	}
	cmds := gjson.GetBytes(data, "cmds").Array()
	taskId := gjson.GetBytes(data, "taskId").String()

	fmt.Println("repoUrl:", repoUrl)
	fmt.Println("serviceName:", serviceName)
	fmt.Println("profile:", Profile)
	fmt.Println("version:", Version)
	fmt.Println("port:", Port)
	fmt.Println("options:", Options)
	fmt.Println("workdir:", workDir)

	fmt.Println("开始遍历cmds...")
	for index := range cmds {
		dir := cmds[index].Get("dir").String()
		cmdStr := cmds[index].Get("cmd").String()

		dir = strings.Replace(dir, "$work_dir", workDir, -1)
		cmdStr = ReplaceStr(cmdStr, repoUrl, serviceName, Profile, Version, Port, Options)
		splits := strings.Split(cmdStr, " ")
		fmt.Println("开始执行命令...")
		if result, err := Excute(splits[0], splits[1:], dir); err != nil {
			returnResult := ReturnResult(index, taskId, -1, "执行第"+cast.ToString(index+1)+"条命令失败")
			fmt.Println("发送结果到消息队列...")
			c.Send <- returnResult

		} else {
			returnResult := ReturnResult(index, taskId, -1, string(result))
			fmt.Println("发送成功的结果到消息队列...")
			c.Send <- returnResult
		}
	}

}

func ReplaceStr(cmdStr, repoUrl, serviceName, Profile, Version, Port, Options string) string {
	if strings.Contains(cmdStr, "$repo_url") {
		cmdStr = strings.Replace(cmdStr, "$repo_url", repoUrl, -1)
	}
	if strings.Contains(cmdStr, "$service_name") {
		cmdStr = strings.Replace(cmdStr, "$service_name", serviceName, -1)
	}

	if strings.Contains(cmdStr, "$profile") {
		cmdStr = strings.Replace(cmdStr, "$profile", Profile, -1)
	}

	if strings.Contains(cmdStr, "$version") {
		cmdStr = strings.Replace(cmdStr, "$version", Version, -1)
	}

	if strings.Contains(cmdStr, "$port") {
		cmdStr = strings.Replace(cmdStr, "$port", Port, -1)
	}

	if strings.Contains(cmdStr, "$options") {
		cmdStr = strings.Replace(cmdStr, "$options", Options, -1)
	}
	return cmdStr
}

func Excute(name string, cmds []string, dir string) (data []byte, err error) {
	var stdout bytes.Buffer

	cmd := exec.Command(name, cmds...)
	cmd.Dir = dir

	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return nil, errors.New("执行命令失败")
	} else {
		return stdout.Bytes(), nil
	}
}

func ReturnResult(index int, taskId string, code int, msg string) []byte {
	var cmdResult models.Result
	cmdResult.Action = "ack"
	cmdResult.SerialId = index + 1
	cmdResult.Cmd = "第" + cast.ToString(index+1) + "条shell命令"
	cmdResult.TaskId = taskId
	cmdResult.Resp.Msg = msg
	cmdResult.Resp.Code = code
	bytes, _ := json.Marshal(cmdResult)
	return bytes
}
