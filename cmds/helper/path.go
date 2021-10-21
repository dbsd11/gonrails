package helper

import (
	"fmt"
	"os"
)

var ProjectPath string  // 生成项目地址
var TemplatePath string // 模版的目录

func init() {
	if "true" == os.Getenv("DEV") {
		ProjectPath, _ = os.Getwd()
		TemplatePath = fmt.Sprintf("%s/cmds/templates", ProjectPath)
	} else {
		ProjectPath, _ = os.Getwd()
		TemplatePath = fmt.Sprintf("%s/pkg/mod/github.com/gonrails/gonrails@v0.1.8/cmds/templates", os.Getenv("GOPATH"))
	}

	fmt.Println(ProjectPath)
	fmt.Println(TemplatePath)
}
