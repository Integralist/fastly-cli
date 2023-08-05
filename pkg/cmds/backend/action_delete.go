package backend

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/global"
	"github.com/integralist/fastly-cli/pkg/helper"
)

// NewDeleteCmd returns a delete action command.
func NewDeleteCmd(g global.Data) *cli.Command {
	c := helper.NewDeleteCmd(Category)
	c.Usage = "Delete a backend"
	c.Action = func(ctx *cli.Context) error {
		fmt.Fprintf(ctx.Command.Writer, "delete backend: '%s'\n", ctx.Args().First())
		return nil
	}
	return c
}
