package backend

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/global"
	"github.com/integralist/fastly-cli/pkg/helper"
	"github.com/integralist/fastly-cli/pkg/manifest"
)

// NewReadCmd returns a create action command.
func NewReadCmd(g global.Data) *cli.Command {
	c := helper.NewReadCmd(Category)
	c.Usage = "Get a backend"
	c.Action = func(ctx *cli.Context) error {
		fmt.Fprintf(ctx.Command.Writer, "get backend: '%s'\n", ctx.Args().First())

		m, err := manifest.Read(manifest.Filename)
		if err != nil {
			return fmt.Errorf("failed to read manifest: %w", err)
		}

		fmt.Printf("\nManifest: %#v\n\n", m)
		fmt.Printf("Config: %#v\n", g.Config)

		return nil
	}
	return c
}
