package snippet

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/global"
	"github.com/integralist/fastly-cli/pkg/helper"
)

// NewCreateCmd returns a create action command.
func NewCreateCmd(g global.Data) *cli.Command {
	c := helper.NewCreateCmd(Category)
	c.Usage = "Create a snippet"
	c.Action = func(ctx *cli.Context) error {
		fmt.Fprintf(ctx.Command.Writer, "create snippet: '%s'\n", ctx.Args().First())
		return nil
	}
	return c
}
