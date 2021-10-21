package cmds

import (
	"fmt"
	"log"
	"os"

	"github.com/gonrails/gonrails/cmds/helper"
	"github.com/gonrails/gonrails/cmds/new"
	"github.com/urfave/cli/v2"
)

func New(ctx *cli.Context) error {
	name := ctx.Args().First()
	makeDirs(name)

	new.TouchMain(name)
	new.TouchConfig(name)
        new.TouchRoute(name)
        new.TouchController(name)
        new.TouchSerializer(name)
        new.TouchDockerfile(name)
        new.InitMod(name)
	new.InitGit(name)

	return nil
}

func makeDirs(name string) {
	log.Printf(helper.Yellow("---> Generating gonrails project named %s ..."), helper.Green(name))
	_ = os.Mkdir(name, os.ModePerm)

	for _, subDir := range dirs {
		log.Println(fmt.Sprintf(helper.Blue("------> Making directory: %s"), fmt.Sprintf(helper.Green("%s/%s"), name, subDir)))
		_ = os.Mkdir(fmt.Sprintf("%s/%s", name, subDir), os.ModePerm)
	}
}

