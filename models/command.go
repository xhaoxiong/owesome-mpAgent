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
	Action   string `json:"action"`
	TaskId   string `json:"taskId"`
	SerialId int    `json:"serialId"`
	Cmd      string `json:"cmd"`
	Resp     struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
}

type HeartBeatResult struct {
	Action string
	Data   struct {
		Ip    string `json:"Ip"`
		Queue string `json:"Queue"`
	}
}
