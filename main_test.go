/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE main
 */

package main

import (
	"github.com/shirou/gopsutil/cpu"
	"testing"
	"time"
)

func TestA(t *testing.T) {

	used, _ := cpu.Percent(2*time.Second, false)

	t.Log(used)

}
