/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE main
 */

package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"mpAgent/config"
	"mpAgent/service"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	t.Log(service.CPUCheck())
	t.Log(service.FreeCPU())
	t.Log(service.DiskCheck())
	t.Log(service.RAMCheck())
	t.Log(viper.GetStringSlice("heart_beat.machine_type"))
	fmt.Println(time.Now().Weekday())
}
