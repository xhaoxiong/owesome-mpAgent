/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE main
 */

package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestA(t *testing.T) {

	var stdout bytes.Buffer
	cmd := exec.Command("ls", "-l")

	cmd.Stdout = &stdout
	cmd.Run()

	fmt.Println(stdout.String())
}
