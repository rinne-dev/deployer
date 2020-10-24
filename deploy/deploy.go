package deploy

import (
	"errors"
	"github.com/codeskyblue/go-sh"
	"github.com/urfave/cli/v2"
	"rinne.dev/deployer/task"
	"strings"
)

// CliConfig 返回 CLI 配置
func CliConfig() *cli.Command {
	return &cli.Command{
		Name: "deploy",
		Aliases: []string{"d"},
		Usage: "部署应用",
		Action: func (ctx *cli.Context) error {
			return Deploy(ctx.Args().Get(0))
		},
	}
}

// Deploy 执行部署操作
func Deploy(appName string) error {
	if appName == "" {
		return errors.New("请指定待部署应用名称")
	}

	err := task.Exist(appName)
	if err != nil {
		return err
	}

	appWorkdir := task.Workdir(appName)
	appCommands := task.Commands(appName)
	runCommands(appWorkdir, appCommands)

	return nil
}

func runCommands(workdir string, commands []string) {
	session := sh.NewSession().SetDir(workdir)
	session.ShowCMD = false
	for _, command := range commands {
		sliced := strings.Split(command, " ")
		session.Command(sliced[0], sliced[1:]).Run()
	}
}
