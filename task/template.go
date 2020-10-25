package task

var cmdTemplate = map[string][]string{
	"git": {
		"git pull",
	},
	"yarn-install": {
		"yarn",
	},
	"yarn-build": {
		"yarn build",
	},
	"npm-install": {
		"npm install",
	},
	"npm-build": {
		"npm run build",
	},
	"hexo-generate": {
		"hexo generate",
	},
}
