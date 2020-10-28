package server

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"rinne.dev/deployer/config"
	"rinne.dev/deployer/deploy"
)

// CliConfig 返回 CLI 配置
func CliConfig() *cli.Command {
	return &cli.Command{
		Name: "serve",
		Aliases: []string{"s"},
		Usage: "启动 HTTP 服务器",
		Action: func (ctx *cli.Context) error {
			return Serve()
		},
	}
}

// Serve 启动 HTTP 服务器
func Serve() error {
	gin.SetMode(config.GetString("env"))
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Deployer@RinNe.Dev",
		})
	})
	r.POST("/:name/deploy.hook", deploy.HTTPHandler)
	return r.Run(config.GetString("listen"))
}
