/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE models
 */
package models

//输入的json字符串结构体
type Command struct {
	Action string      `json:"action"`
	TaskId string      `json:"taskId"`
	Params interface{} `json:"params"`
	Cmds   []struct {
		Dir string `json:"dir"`
		Cmd string `json:"cmd"`
	}
}

//返回的json字符串结构体
type Result struct {
	Action string `json:"action"`
	Resp   struct {
		Code      int    `json:"code"`
		Msg       string `json:"msg"`
		TaskId    string `json:"taskId"`
		SerialNo  int    `json:"serialId"`
		Cmd       string `json:"cmd"`
		TimeStamp int64  `json:"timestamp"`
	}
}

type HeartBeatResult struct {
	Action string
	Data   struct {
		Ip          string   `json:"Ip"`
		Queue       string   `json:"Queue"`
		Alias       string   `json:"alias"`
		User        string   `json:"user"`
		Passwd      string   `json:"passwd"`
		FreeCpu     int64    `json:"freeCpu"`
		FreeMem     int64    `json:"freeMem"`
		FreeDisk    int64    `json:"freeDisk"`
		MachineType []string `json:"machineType"`
	}
}
