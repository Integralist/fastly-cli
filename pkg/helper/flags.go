package helper

import (
	"fmt"
	"sort"

	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/env"
	"github.com/integralist/fastly-cli/pkg/global"
)

// MergeFlagsWithGlobal merges command local flags with global flags.
// This is necessary as we can't create global flags with urfave/cli.
// So we provide a helper to make it easy to expose flags defined as global.
func MergeFlagsWithGlobal(flags []cli.Flag) []cli.Flag {
	var combined []cli.Flag
	combined = append(combined, flags...)
	combined = append(combined, global.Flags...)
	sort.Sort(cli.FlagsByName(combined))
	return combined
}

// FlagAutoClone is a commonly used flag.
var FlagAutoClone = &cli.BoolFlag{
	Name:  "autoclone",
	Usage: "If the selected service version is not editable, clone it and use the clone",
}

// FlagCustomerID is a commonly used flag.
var FlagCustomerID = &cli.StringFlag{
	Name:  "customer-id",
	Usage: fmt.Sprintf("Alphanumeric string identifying the customer (falls back to %s)", env.CustomerID),
}

// FlagServiceID is a commonly used flag.
var FlagServiceID = &cli.StringFlag{
	Name:  "service-id",
	Usage: fmt.Sprintf("Service ID (falls back to %s, then fastly.toml)", env.ServiceID),
}

// FlagServiceName is a commonly used flag.
var FlagServiceName = &cli.StringFlag{
	Name:  "service-name",
	Usage: "The name of the service",
}

// FlagServiceVersion is a commonly used flag.
var FlagServiceVersion = &cli.StringFlag{
	Name:  "service-version",
	Usage: "'latest', 'active', or the number of a specific Fastly service version",
}
