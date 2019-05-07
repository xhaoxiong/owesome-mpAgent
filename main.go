/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE mpAgent
 */
package main

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/pflag"
	"github.com/tidwall/gjson"
	"mpAgent/config"
	"mpAgent/service"
)

var (
	cfg = pflag.StringP("config", "c", "", "./config.yaml")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	cmd := service.NewCommand()
	mq := service.NewRabbitMq()
	go service.HeartBeatTicker(mq)
	go func() {
		for {

			if data := mq.Consumer(); data != nil {
				fmt.Println(string(data))
				var t interface{}
				gjson.Unmarshal(data, &t)
				data = []byte(cast.ToString(t))
				cmd.Recv <- data
				cmd.Start()
			}
		}
	}()

	go func() {
		for {
			if data, ok := <-cmd.Send; ok {
				fmt.Println(string(data))
				if data != nil {
					mq.Publish(data)
				}
			}
		}
	}()

	select {}
}

func A() {

}
