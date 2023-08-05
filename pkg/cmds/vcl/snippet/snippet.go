// Package snippet contains sub commands for managing versioned and dynamic VCL.
package snippet

import (
	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/global"
)

// Category is the subcommand category.
const Category = "VCL"

// NewCmd returns a category command.
func NewCmd(g global.Data) *cli.Command {
	return &cli.Command{
		Name:     "snippet",
		Aliases:  []string{"s"},
		Usage:    "Manage VCL snippets",
		Category: Category,
		Commands: []*cli.Command{
			NewCreateCmd(g),
			NewDeleteCmd(g),
			NewListCmd(g),
			NewReadCmd(g),
			NewUpdateCmd(g),
		},
	}
}
