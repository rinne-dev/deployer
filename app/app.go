package app

import (
	"github.com/urfave/cli/v2"
	"rinne.dev/deployer/config"
	"rinne.dev/deployer/deploy"
	"rinne.dev/deployer/server"
)

const VERSION = "v0.1.1"

// App 返回 CLI 实例
func App() *cli.App {
	return &cli.App{
		Name: "Deployer@RinNe.Dev",
		Usage: "一个简单易用的部署工具",
		Version: VERSION,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				Aliases: []string{"c"},
				Value: "config.yaml",
				Usage: "Load configuration from `FILE`",
			},
		},
		Before: func(ctx *cli.Context) error {
			return config.Load(ctx.String("config"))
		},
		Commands: []*cli.Command{
			server.CliConfig(),
			deploy.CliConfig(),
		},
	}
}
