/**
*@Author: haoxiongxiao
*@Date: 2019/5/6
*@Description: CREATE GO FILE service
 */
package service

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"time"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// DiskCheck checks the disk usage.
func DiskCheck() (message string) {
	u, _ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	text := "OK"

	if usedPercent >= 95 {

		text = "CRITICAL"
	} else if usedPercent >= 90 {

		text = "WARNING"
	}

	message = fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	return
}

func FreeDisk() int {
	u, _ := disk.Usage("/")
	return int(u.Free) / MB
}

func TotalDisk() int {
	u, _ := disk.Usage("/")
	return int(u.Total) / MB
}

// CPUCheck checks the cpu usage.
func CPUCheck() (message string) {
	cores, _ := cpu.Counts(false)
	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	text := "OK"

	if l5 >= float64(cores-1) {

		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		text = "WARNING"
	}
	used, _ := cpu.Percent(time.Second, false)
	message = fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f |Cores: %d |Used: %.2f%%", text, l1, l5, l15, cores, used[0])
	return
}

func FreeCPU() float64 {
	used, _ := cpu.Percent(time.Second, false)

	return used[0]
}

func CPUNum() int {
	cores, _ := cpu.Counts(false)
	return cores
}

// RAMCheck checks the disk usage.
func RAMCheck() (message string) {
	u, _ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	text := "OK"

	if usedPercent >= 95 {

		text = "CRITICAL"
	} else if usedPercent >= 90 {

		text = "WARNING"
	}

	message = fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	return
}

func FreeRAM() int {
	u, _ := mem.VirtualMemory()
	return int(u.Free) / MB
}

func TotalRAM() int {
	u, _ := mem.VirtualMemory()
	return int(u.Total) / MB

}
