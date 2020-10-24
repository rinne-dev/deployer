package webhook

import (
	"gopkg.in/go-playground/webhooks.v5/github"
	"net/http"
)

// Github 解析 GitHub Webhook 请求
func Github(req *http.Request, secret string) error {
	hook, err := github.New(github.Options.Secret(secret))
	if err != nil {
		return err
	}

	_, err = hook.Parse(req, github.PushEvent)
	if err != nil {
		return err
	}

	return nil
}
