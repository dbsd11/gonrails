package cmds

import (
	"github.com/gonrails/gonrails/cmds/generate/controller"
	"github.com/gonrails/gonrails/cmds/generate/migration"
        "github.com/urfave/cli/v2"
)

// Generate - for command generate
func Generate(ctx *cli.Context) error {
        args := ctx.Args().Slice()
	if "controller" == args[1] {
		controller.Exec(args)
	}

	if "migration" == args[1] {
		migration.Exec(args)
	}

	if "model" == args[1] {

	}

        return nil
}
