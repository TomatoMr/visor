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
	}
}

func record(per float64) {
	now := time.Now().Format("2006-01-02-15-04-05")
	logFile := config.GetConfig().SnapPath + now + ".log"
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
