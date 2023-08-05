// Package vcl provides sub commands for managing VCL.
package vcl

import (
	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/cmds/vcl/snippet"
	"github.com/integralist/fastly-cli/pkg/global"
)

// NewCmd returns a category command.
func NewCmd(g global.Data) *cli.Command {
	return &cli.Command{
		Name:     "vcl",
		Aliases:  []string{"v"},
		Usage:    "Manage VCL",
		Category: "VCL",
		Commands: []*cli.Command{
			snippet.NewCmd(g),
		},
	}
}
