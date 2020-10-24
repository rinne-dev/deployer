package task

import (
	"errors"
	"rinne.dev/deployer/config"
)

// Exist 判断任务是否存在
func Exist(appName string) error {
	if config.Get("tasks." + appName) != nil {
		return nil
	}
	return errors.New("找不到此应用")
}

// Type 获取任务类型
func Type(appName string) string {
	return config.GetString("tasks." + appName + ".type")
}

// Secret 获取任务 Secret
func Secret(appName string) string {
	return config.GetString("tasks." + appName + ".secret")
}

// Workdir 获取任务工作目录
func Workdir(appName string) string {
	return config.GetString("tasks." + appName + ".workdir")
}

// Templates 获取任务部署命令模板
func Templates(appName string) []string {
	return config.GetStringSlice("tasks." + appName + ".templates")
}

// Commands 获取任务部署命令
func Commands(appName string) []string {
	var commands []string

	templates := Templates(appName)
	for _, template := range templates {
		if template == "custom" {
			custom := config.GetStringSlice("tasks." + appName + ".commands")
			commands = append(commands, custom...)
			continue
		}

		if cmd, ok := cmdTemplate[template]; ok {
			commands = append(commands, cmd...)
		}
	}

	return commands
}
