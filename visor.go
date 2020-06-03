package main

import (
	"github.com/TomatoMr/visor/config"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"gopkg.in/gomail.v2"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func visorCpu() {
	for {
		per, _ := cpu.Percent(time.Duration(config.GetConfig().Interval)*time.Second, false)
		if per[0] > config.GetConfig().AlterLimit {
			record()
		}
	}
}

func visorMem() {
	for {
		m, _ := mem.VirtualMemory()
		if m.UsedPercent > config.GetConfig().AlterLimit {
			record()
		}
		time.Sleep(time.Duration(config.GetConfig().Interval) * time.Second)
	}
}

func record() {
	now := time.Now().Format("2006-01-02-15-04-05")
	logFile := config.GetConfig().SnapPath + now + ".log"
	f, _ := os.Create(logFile)
	defer f.Close()
	//loadCmd := exec.Command("w")
	//loadOutput, _ := loadCmd.Output()

	f.Write([]byte{'\n'})
	f.Write([]byte{'\n'})
	topCmd := exec.Command("top", "-H", "-w", "512", "-c", "-n", "1", "-b")
	topOutput, _ := topCmd.Output()
	f.Write(topOutput)
	SendMail("你服务器炸了", string(topOutput), config.GetConfig())
}

func SendMail(subject string, body string, conf config.Config) error {
	mailConn := map[string]string{
		"user": conf.FromMail,
		"pass": conf.FromMailPass,
		"host": conf.FromMailHost,
		"port": conf.FromMailPort,
	}

	port, _ := strconv.Atoi(mailConn["port"])

	m := gomail.NewMessage()
	m.SetHeader("From", "<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", conf.ToMail...)             //发送给多个用户
	m.SetHeader("Subject", subject)               //设置邮件主题
	m.SetBody("text/html", body)                  //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err
}
