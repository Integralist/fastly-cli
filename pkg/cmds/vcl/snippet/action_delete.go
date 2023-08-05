package snippet

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/global"
	"github.com/integralist/fastly-cli/pkg/helper"
)

// NewDeleteCmd returns a create action command.
func NewDeleteCmd(g global.Data) *cli.Command {
	c := helper.NewDeleteCmd(Category)
	c.Usage = "Delete a snippet"
	c.Action = func(ctx *cli.Context) error {
		fmt.Fprintf(ctx.Command.Writer, "delete snippet: '%s'\n", ctx.Args().First())
		return nil
	}
	return c
}
