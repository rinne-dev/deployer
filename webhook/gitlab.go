package webhook

import (
	"gopkg.in/go-playground/webhooks.v5/gitlab"
	"net/http"
)

// Gitlab 解析 Gitlab Webhook 请求
func Gitlab(req *http.Request, secret string) error {
	hook, err := gitlab.New(gitlab.Options.Secret(secret))
	if err != nil {
		return err
	}

	_, err = hook.Parse(req, gitlab.PushEvents)
	if err != nil {
		return err
	}

	return nil
}
