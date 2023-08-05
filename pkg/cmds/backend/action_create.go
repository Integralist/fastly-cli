package backend

import (
	"fmt"

	"github.com/fastly/go-fastly/v8/fastly"
	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/global"
	"github.com/integralist/fastly-cli/pkg/helper"
)

// NewCreateCmd returns a create action command.
func NewCreateCmd(g global.Data) *cli.Command {
	c := helper.NewCreateCmd(Category)
	c.Usage = "Create a backend"
	c.Flags = flagsForCreate()
	c.Action = func(ctx *cli.Context) error {
		sid, sv, err := helper.ServiceDetails(ctx, g)
		if err != nil {
			return err
		}

		v, err := g.APIClient.CreateBackend(opts(ctx, sid, sv))
		if err != nil {
			return err
		}

		fmt.Fprintf(ctx.Command.Writer, "Created backend '%s' (service '%s' version '%d')", v.Name, v.ServiceID, v.ServiceVersion)
		return nil
	}
	return c
}

func flagsForCreate() []cli.Flag {
	helper.FlagServiceVersion.Required = true

	return helper.MergeFlagsWithGlobal([]cli.Flag{
		helper.FlagAutoClone,
		helper.FlagServiceID,
		helper.FlagServiceName,
		helper.FlagServiceVersion,
		&cli.StringFlag{
			Name:  "address",
			Usage: "A hostname, IPv4, or IPv6 address for the backend",
		},
		&cli.BoolFlag{
			Name:  "auto-loadbalance",
			Usage: "Whether or not this backend should be automatically load balanced",
		},
		&cli.IntFlag{
			Name:  "between-bytes-timeout",
			Usage: "How long to wait between bytes in milliseconds",
		},
		&cli.StringFlag{
			Name:  "comment",
			Usage: "A descriptive note",
		},
		&cli.IntFlag{
			Name:  "connect-timeout",
			Usage: "How long to wait for a timeout in milliseconds",
		},
		&cli.IntFlag{
			Name:  "first-byte-timeout",
			Usage: "How long to wait for the first bytes in milliseconds",
		},
		&cli.StringFlag{
			Name:  "healthcheck",
			Usage: "The name of the healthcheck to use with this backend",
		},
		&cli.IntFlag{
			Name:  "keep-alive-time",
			Usage: "How long in seconds to keep a persistent connection to the backend between requests",
		},
		&cli.IntFlag{
			Name:  "max-conn",
			Usage: "Maximum number of connections",
		},
		&cli.StringFlag{
			Name:  "max-tls-version",
			Usage: "Maximum allowed TLS version on SSL connections to this backend",
		},
		&cli.StringFlag{
			Name:  "min-tls-version",
			Usage: "Minimum allowed TLS version on SSL connections to this backend",
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "Backend name",
		},
		&cli.StringFlag{
			Name:  "override-host",
			Usage: "The hostname to override the Host header",
		},
		&cli.IntFlag{
			Name:  "port",
			Usage: "Port number of the address",
		},
		&cli.StringFlag{
			Name:  "request-condition",
			Usage: "Condition, which if met, will select this backend during a request",
		},
		&cli.StringFlag{
			Name:  "shield",
			Usage: "The shield POP designated to reduce inbound load on this origin by serving the cached data to the rest of the network",
		},
		&cli.StringFlag{
			Name:  "ssl-ca-cert",
			Usage: "CA certificate attached to origin",
		},
		&cli.StringFlag{
			Name:  "ssl-cert-hostname",
			Usage: "Overrides ssl_hostname, but only for cert verification. Does not affect SNI at all",
		},
		&cli.BoolFlag{
			Name:  "ssl-check-cert",
			Usage: "Be strict on checking SSL certs",
		},
		&cli.StringFlag{
			Name:  "ssl-ciphers",
			Usage: "List of OpenSSL ciphers (https://www.openssl.org/docs/man1.0.2/man1/ciphers)",
		},
		&cli.StringFlag{
			Name:  "ssl-client-cert",
			Usage: "Client certificate attached to origin",
		},
		&cli.StringFlag{
			Name:  "ssl-client-key",
			Usage: "Client key attached to origin",
		},
		&cli.StringFlag{
			Name:  "ssl-sni-hostname",
			Usage: "Overrides ssl_hostname, but only for SNI in the handshake. Does not affect cert validation at all",
		},
		&cli.BoolFlag{
			Name:  "use-ssl",
			Usage: "Whether or not to use SSL to reach the backend",
		},
		&cli.IntFlag{
			Name:  "weight",
			Usage: "Weight used to load balance this backend against others",
		},
	})
}

func opts(ctx *cli.Context, sid string, sv int) *fastly.CreateBackendInput {
	input := &fastly.CreateBackendInput{}
	input.ServiceID = sid
	input.ServiceVersion = sv

	if ctx.IsSet("address") {
		input.Address = fastly.String(ctx.Value("address").(string))
	}
	if ctx.IsSet("auto-loadbalance") {
		input.AutoLoadbalance = fastly.CBool(ctx.Value("auto-loadbalance").(bool))
	}
	if ctx.IsSet("between-bytes-timeout") {
		input.BetweenBytesTimeout = fastly.Int(int(ctx.Value("between-bytes-timeout").(int64)))
	}
	if ctx.IsSet("comment") {
		input.Comment = fastly.String(ctx.Value("comment").(string))
	}
	if ctx.IsSet("connect-timeout") {
		input.ConnectTimeout = fastly.Int(int(ctx.Value("connect-timeout").(int64)))
	}
	if ctx.IsSet("first-byte-timeout") {
		input.FirstByteTimeout = fastly.Int(int(ctx.Value("first-byte-timeout").(int64)))
	}
	if ctx.IsSet("healthcheck") {
		input.HealthCheck = fastly.String(ctx.Value("healthcheck").(string))
	}
	if ctx.IsSet("keep-alive-time") {
		input.KeepAliveTime = fastly.Int(int(ctx.Value("keep-alive-time").(int64)))
	}
	if ctx.IsSet("max-conn") {
		input.FirstByteTimeout = fastly.Int(int(ctx.Value("max-conn").(int64)))
	}
	if ctx.IsSet("max-tls-version") {
		input.MaxTLSVersion = fastly.String(ctx.Value("max-tls-version").(string))
	}
	if ctx.IsSet("min-tls-version") {
		input.MinTLSVersion = fastly.String(ctx.Value("min-tls-version").(string))
	}
	if ctx.IsSet("name") {
		input.Name = fastly.String(ctx.Value("name").(string))
	}
	if ctx.IsSet("override-host") {
		input.OverrideHost = fastly.String(ctx.Value("override-host").(string))
	}
	if ctx.IsSet("port") {
		input.Port = fastly.Int(int(ctx.Value("port").(int64)))
	}
	if ctx.IsSet("request-condition") {
		input.RequestCondition = fastly.String(ctx.Value("request-condition").(string))
	}
	if ctx.IsSet("shield") {
		input.Shield = fastly.String(ctx.Value("shield").(string))
	}
	if ctx.IsSet("ssl-ca-cert") {
		input.SSLCACert = fastly.String(ctx.Value("ssl-ca-cert").(string))
	}
	if ctx.IsSet("ssl-cert-hostname") {
		input.SSLCertHostname = fastly.String(ctx.Value("ssl-cert-hostname").(string))
	}
	if ctx.IsSet("ssl-check-cert") {
		input.SSLCheckCert = fastly.CBool(ctx.Value("ssl-check-cert").(bool))
	}
	if ctx.IsSet("ssl-ciphers") {
		input.SSLCiphers = fastly.String(ctx.Value("ssl-ciphers").(string))
	}
	if ctx.IsSet("ssl-client-cert") {
		input.SSLClientCert = fastly.String(ctx.Value("ssl-client-cert").(string))
	}
	if ctx.IsSet("ssl-client-key") {
		input.SSLClientKey = fastly.String(ctx.Value("ssl-client-key").(string))
	}
	if ctx.IsSet("ssl-sni-hostname") {
		input.SSLSNIHostname = fastly.String(ctx.Value("ssl-sni-hostname").(string))
	}
	if ctx.IsSet("use-ssl") {
		input.UseSSL = fastly.CBool(ctx.Value("use-ssl").(bool))
	}
	if ctx.IsSet("weight") {
		input.Weight = fastly.Int(int(ctx.Value("weight").(int64)))
	}

	return input
}
