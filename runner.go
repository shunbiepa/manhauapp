package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"manhauapp/config"
	"manhauapp/global"
	"manhauapp/middleware/cache"
	"time"

	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func InitServer() {
	var err error
	global.GVA_CONFIG, err = config.Get(configPath) // 初始化配置
	if err != nil {
		log.Fatalln("read conf file : ", err)
	}

	global.GVA_LOG = myzap.Zap() // 初始化zap日志库

	global.GVA_DB, err = db.Instance(global.GVA_CONFIG) // 初始化数据库库
	if err != nil {
		log.Fatalln("init db err : ", err)
	}

	global.GVA_REDIS, err = cache.InitRedis(global.GVA_CONFIG.Cache.Addr, global.GVA_CONFIG.Cache.Password, 0) // 初始化redis服务
	if err != nil {
		log.Fatalln("redis init err : ", err)
	}

	auth.InitPermission() // 初始化权限

	service.DataUploadHandler()
}

// StartServer 启动服务
func StartServer() {

	// 初始化路由
	Router := router.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.Port)
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	global.GVA_LOG.Error(s.ListenAndServe().Error())

}

// initServer 借助endless 服务的实现优雅重启
func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 60 * time.Millisecond
	s.WriteTimeout = 60 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
