package webhook

import (
	"gopkg.in/go-playground/webhooks.v5/gogs"
	"net/http"
)

// Gogs 解析 Gogs Webhook 请求
func Gogs(req *http.Request, secret string) error {
	hook, err := gogs.New(gogs.Options.Secret(secret))
	if err != nil {
		return err
	}

	_, err = hook.Parse(req, gogs.PushEvent)
	if err != nil {
		return err
	}

	return nil
}
