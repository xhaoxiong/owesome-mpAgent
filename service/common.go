/**
*@Package:service
*@Author: haoxiongxiao
*@Date: 2019/5/7
*@Description: create go file in service package
 */
package service

import (
	"bytes"
	"encoding/json"
	"mpAgent/models"
	"os"
	"os/exec"
	"strings"
	"time"
)

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

func cmdReplace(cmdStr, repoUrl, serviceName, Profile, Version, Port, Options, workDir string) string {
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

func optionsReplace(opt, port string) (options string) {
	options = strings.Replace(opt, "$port", port, -1)
	options = strings.Replace(options, "$HOST", os.Getenv("HOST"), -1)
	options = strings.Replace(options, "$REGISTER_CENTER", os.Getenv("REGISTER_CENTER"), -1)
	return
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
