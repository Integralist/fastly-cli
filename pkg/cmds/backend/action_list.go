package backend

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/global"
	"github.com/integralist/fastly-cli/pkg/helper"
)

// NewListCmd returns a delete action command.
func NewListCmd(g global.Data) *cli.Command {
	c := helper.NewListCmd(Category)
	c.Usage = "List backends"
	c.Action = func(ctx *cli.Context) error {
		fmt.Fprintf(ctx.Command.Writer, "list backend: '%s'\n", ctx.Args().First())
		return nil
	}
	return c
}
