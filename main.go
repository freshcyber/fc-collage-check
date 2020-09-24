package main

import (
	"fmt"

	"github.com/robfig/cron"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/freshcyber/fc-collage-check/config"
	"github.com/freshcyber/fc-collage-check/handler"
	"github.com/freshcyber/fc-collage-check/persist"
)

func main() {

	persist.GMariadb.Init()

	fmt.Println("Init MariaDB")

	c := cron.New()

	// 查询拼单
	c.AddJob(config.Server.Spec, handler.CollageOrderInfo{})

	c.Start()

	fmt.Println("==== start jobs ====")

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select {}
}
