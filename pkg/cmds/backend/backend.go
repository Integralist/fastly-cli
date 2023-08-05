// Package backend provides sub commands for managing backends (origin servers).
package backend

import (
	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/global"
)

// Category is the subcommand category.
const Category = "backend"

// NewCmd returns a category command.
func NewCmd(g global.Data) *cli.Command {
	return &cli.Command{
		Name:     "backend",
		Aliases:  []string{"b"},
		Usage:    "Manage service backends",
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
