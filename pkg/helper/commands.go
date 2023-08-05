package helper

import "github.com/urfave/cli/v3"

// NewCreateCmd returns a create action command with predefined settings.
func NewCreateCmd(category string) *cli.Command {
	return &cli.Command{
		Name:     "create",
		Aliases:  []string{"c", "add", "a"},
		Category: category,
	}
}

// NewDeleteCmd returns a delete action command with predefined settings.
func NewDeleteCmd(category string) *cli.Command {
	return &cli.Command{
		Name:     "delete",
		Aliases:  []string{"d", "remove", "r"},
		Category: category,
	}
}

// NewListCmd returns a read action command with predefined settings.
func NewListCmd(category string) *cli.Command {
	return &cli.Command{
		Name:     "list",
		Aliases:  []string{"l"},
		Category: category,
	}
}

// NewReadCmd returns a read action command with predefined settings.
func NewReadCmd(category string) *cli.Command {
	return &cli.Command{
		Name:     "describe",
		Aliases:  []string{"d", "get", "g", "read", "r"},
		Category: category,
	}
}

// NewUpdateCmd returns a update action command with predefined settings.
func NewUpdateCmd(category string) *cli.Command {
	return &cli.Command{
		Name:     "update",
		Aliases:  []string{"u"},
		Category: category,
	}
}
