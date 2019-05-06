/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE mpAgent
 */
package main

import (
	"fmt"
	"github.com/spf13/pflag"
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

	mq.Publish(service.HeartBeat())

	go func() {
		for {
			if data := mq.Consumer(); data != nil {
				fmt.Println(string(data))
				cmd.Recv <- data
				cmd.Start()
			}
		}
	}()

	go func() {
		for {
			if data, ok := <-cmd.Send; ok {
				fmt.Println(string(data))
				mq.Publish(data)
			}
		}
	}()

	select {}
}

func A() {

}
