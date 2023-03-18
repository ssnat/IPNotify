package main

import (
	"fmt"
	lib "github.com/PxGo/IPNotify/lib"
	"github.com/robfig/cron/v3"
)

func main() {

	var originIP = ""

	config := lib.GetConfig()

	crontab := cron.New(cron.WithSeconds())

	task := func() {

		lib.Logger.Info("Starting scheduled task...")

		ip, err := lib.GetIPv4()
		if err != nil {
			lib.Logger.Error(err)
			return
		}

		if ip == originIP {
			lib.Logger.Info("IP address remains unchanged.")
			lib.Logger.Info("Scheduled task has completed.")
			return
		}

		if originIP == "" {
			err = lib.SendStartupEmail(ip)
		} else {
			err = lib.SendIPChangeEmail(ip, originIP)
		}

		if err != nil {
			lib.Logger.Error(err)
			return
		}
		originIP = ip

		lib.Logger.Info("Scheduled task has completed.")
	}

	_, err := crontab.AddFunc(config.Ip.Interval, task)
	if err != nil {
		lib.Logger.PanicError(err)
	}
	crontab.Start()
	lib.Logger.Info(fmt.Sprintf("Scheduled task [%s] has been started.", config.Ip.Interval))

	select {}
}
