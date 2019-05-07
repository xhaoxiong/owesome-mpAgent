/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE service
 */
package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"github.com/xhaoxiong/util"
	"mpAgent/models"
	"os"
	"os/exec"
	"strings"
	"time"
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
	Options = strings.Replace(Options, "$port", Port, -1)

	if !util.Exists(workDir) {
		os.MkdirAll(workDir, 0755)
	}
	cmds := gjson.GetBytes(data, "cmds").Array()
	taskId := gjson.GetBytes(data, "taskId").String()

	fmt.Println("开始遍历cmds...")
	for index := range cmds {
		dir := cmds[index].Get("dir").String()
		cmdStr := cmds[index].Get("cmd").String()

		dir = strings.Replace(dir, "$work_dir", workDir, -1)
		dir = strings.Replace(dir, "$service_name", serviceName, -1)
		if !util.Exists(dir) {
			os.MkdirAll(dir, 0755)
		}
		cmdStr = ReplaceStr(cmdStr, repoUrl, serviceName, Profile, Version, Port, Options, workDir)
		splits := strings.Split(cmdStr, " ")
		if result, err := Excute(splits[0], splits[1:], dir); err != nil {
			returnResult := ReturnResult(index, taskId, -1, cast.ToString(err), cmdStr)

			c.Send <- returnResult

		} else {
			returnResult := ReturnResult(index, taskId, 0, string(result), cmdStr)
			c.Send <- returnResult
		}
	}
	fmt.Println("遍历cmds完成...")
	c.Send <- ReturnEnd(taskId)

}
func Excute(name string, cmds []string, dir string) (data []byte, err error) {
	var stdout bytes.Buffer

	cmd := exec.Command(name, cmds...)
	cmd.Dir = dir

	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return stdout.Bytes(), nil
}

func ReplaceStr(cmdStr, repoUrl, serviceName, Profile, Version, Port, Options, workDir string) string {
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

	if strings.Contains(cmdStr, "$work_dir") {
		cmdStr = strings.Replace(cmdStr, "$work_dir", workDir, -1)
	}
	return cmdStr
}

func ReturnResult(index int, taskId string, code int, msg string, cmdStr string) []byte {
	var cmdResult models.Result
	cmdResult.Action = "ack"
	cmdResult.Data.SerialNo = index + 1
	cmdResult.Data.Cmd = cmdStr
	cmdResult.Data.TaskId = taskId
	cmdResult.Data.Msg = msg
	cmdResult.Data.Code = code
	cmdResult.Data.TimeStamp = time.Now().Unix()
	bytes, _ := json.Marshal(cmdResult)
	return bytes
}

func ReturnEnd(taskId string) []byte {
	var cmdResult models.Result
	cmdResult.Data.Code = 10000
	cmdResult.Data.TimeStamp = time.Now().Unix()
	cmdResult.Data.TaskId = taskId
	cmdResult.Action = "ack"
	bytes, _ := json.Marshal(cmdResult)
	return bytes
}
