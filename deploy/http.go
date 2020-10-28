package deploy

import (
	"github.com/gin-gonic/gin"
	"log"
	"rinne.dev/deployer/task"
	"rinne.dev/deployer/webhook"
)

// HTTPHandler HTTP 处理器
func HTTPHandler(c *gin.Context) {
	appName := c.Param("name")
	err := task.Exist(appName)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "app not found",
		})
		return
	}

	err = nil
	appType := task.Type(appName)
	appSecret := task.Secret(appName)
	switch appType {
	case "github":
		err = webhook.Github(c.Request, appSecret)
		break
	case "gogs":
		err = webhook.Gogs(c.Request, appSecret)
		break
	case "gitlab":
		err = webhook.Gitlab(c.Request, appSecret)
		break
	default:
		c.JSON(400, gin.H{
			"message": "not supported",
		})
		return
	}
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "verification failed",
		})
		return
	}

	go func() {
		err := Deploy(appName)
		if err != nil {
			log.Println(err)
		}
	}()

	c.JSON(200, gin.H{
		"message": "ok",
	})
}
