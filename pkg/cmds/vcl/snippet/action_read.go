package snippet

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/global"
	"github.com/integralist/fastly-cli/pkg/helper"
)

// NewReadCmd returns a create action command.
func NewReadCmd(g global.Data) *cli.Command {
	c := helper.NewReadCmd(Category)
	c.Usage = "Get a snippet"
	c.Action = func(ctx *cli.Context) error {
		fmt.Fprintf(ctx.Command.Writer, "get snippet: '%s'\n", ctx.Args().First())
		return nil
	}
	return c
}
