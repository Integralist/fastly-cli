package snippet

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/global"
	"github.com/integralist/fastly-cli/pkg/helper"
)

// NewListCmd returns a create action command.
func NewListCmd(g global.Data) *cli.Command {
	c := helper.NewListCmd(Category)
	c.Usage = "List snippets"
	c.Action = func(ctx *cli.Context) error {
		fmt.Fprintf(ctx.Command.Writer, "list snippet: '%s'\n", ctx.Args().First())
		return nil
	}
	return c
}
