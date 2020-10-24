package task

var cmdTemplate = map[string][]string{
	"git": {
		"git pull",
	},
	"yarn-install-build": {
		"yarn",
		"yarn build",
	},
	"npm-install-build": {
		"npm install",
		"npm run build",
	},
	"hexo-generate": {
		"hexo generate",
	},
}
