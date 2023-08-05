package snippet

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/global"
	"github.com/integralist/fastly-cli/pkg/helper"
)

// NewUpdateCmd returns a create action command.
func NewUpdateCmd(g global.Data) *cli.Command {
	c := helper.NewUpdateCmd(Category)
	c.Usage = "Update a snippet"
	c.Category = "VCL"
	c.Action = func(ctx *cli.Context) error {
		fmt.Fprintf(ctx.Command.Writer, "update snippet: '%s'\n", ctx.Args().First())
		return nil
	}
	return c
}
