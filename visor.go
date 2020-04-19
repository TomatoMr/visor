package main

import (
	"github.com/TomatoMr/visor/config"
	"github.com/shirou/gopsutil/cpu"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func visor() {
	for {
		per, _ := cpu.Percent(time.Duration(config.GetConfig().Interval)*time.Second, false)
		if per[0] > config.GetConfig().AlterLimit {
			record(per[0])
		}
		time.Sleep(time.Duration(config.GetConfig().Interval) * time.Second)
	}
}

func record(per float64) {
	now := time.Now().Format(time.RFC3339)
	logFile := now + ".log"
	f, _ := os.Create(logFile)
	defer f.Close()
	//loadCmd := exec.Command("w")
	//loadOutput, _ := loadCmd.Output()

	f.Write([]byte{'\n'})
	f.Write([]byte("此时cpu占用率：" + strconv.FormatFloat(per, 'f', 10, 64)))
	f.Write([]byte{'\n'})
	topCmd := exec.Command("top", "-H", "-n", "1", "-b")
	topOutput, _ := topCmd.Output()
	f.Write(topOutput)
}
