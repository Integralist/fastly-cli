// Package app provides a single entry point for the CLI.
package app

import (
	"context"
	"io"

	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/cmds/backend"
	"github.com/integralist/fastly-cli/pkg/cmds/vcl"
	"github.com/integralist/fastly-cli/pkg/global"
)

func Run(args []string, stdin io.Reader, stdout io.Writer) error {
	c := &cli.Command{
		Name:   "fastly",
		Usage:  "A CLI for interacting with the Fastly platform (https://developer.fastly.com/reference/cli/)",
		Reader: stdin,
		Writer: stdout,
	}

	c.Commands = []*cli.Command{
		backend.NewCmd(global.Container),
		vcl.NewCmd(global.Container),
	}

	return c.Run(context.Background(), args)
}
