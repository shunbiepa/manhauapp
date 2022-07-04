package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

var configPath string

func main() {

	flag.StringVar(&configPath, "config", "./conf/conf.yaml", "config file path")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":   "www.flysnow.org",
			"wechat": "flysnow_org",
		})
	})
	r.Run(":8080")
}
