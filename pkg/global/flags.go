package global

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/env"
)

// Flags represents global flags.
var Flags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "accept-defaults",
		Aliases: []string{"d"},
		Usage:   "Accept default options for all interactive prompts apart from Yes/No confirmations",
	},
	&cli.BoolFlag{
		Name:    "auto-yes",
		Aliases: []string{"y"},
		Usage:   "Answer yes automatically to all Yes/No confirmations. This may suppress security warnings",
	},
	&cli.StringFlag{
		Name:        "endpoint",
		Usage:       "Fastly API endpoint",
		DefaultText: "https://api.fastly.com",
	},
	&cli.BoolFlag{
		Name:    "json",
		Aliases: []string{"j"},
		Usage:   "Render output as JSON",
	},
	&cli.BoolFlag{
		Name:    "non-interactive",
		Aliases: []string{"i"},
		Usage:   "Do not prompt for user input - suitable for CI processes. Equivalent to --accept-defaults and --auto-yes",
	},
	&cli.StringFlag{
		Name:    "profile",
		Aliases: []string{"o"},
		Usage:   "Switch account profile for single command execution (see also: 'fastly profile switch')",
	},
	&cli.BoolFlag{
		Name:    "quiet",
		Aliases: []string{"q"},
		Usage:   "Silence all output except direct command output. This won't prevent interactive prompts (see: --accept-defaults, --auto-yes, --non-interactive)",
	},
	&cli.StringFlag{
		Name:    "token",
		Aliases: []string{"t"},
		Usage:   fmt.Sprintf("Fastly API token (or via %s)", env.Token),
	},
	&cli.BoolFlag{
		Name:    "verbose",
		Aliases: []string{"v"},
		Usage:   "Verbose logging",
	},
}
