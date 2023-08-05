package helper

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/fastly/go-fastly/v8/fastly"
	"github.com/urfave/cli/v3"

	"github.com/integralist/fastly-cli/pkg/env"
	"github.com/integralist/fastly-cli/pkg/global"
)

// ServiceDetails returns the Service ID and the version that should be used.
func ServiceDetails(ctx *cli.Context, g global.Data) (sid string, sv int, err error) {
	if sid, err = ServiceID(ctx, g); err != nil {
		return sid, sv, err
	}
	var v *fastly.Version
	if v, err = ServiceVersion(ctx, g, sid); err != nil {
		return sid, sv, err
	}
	if sv, err = AutoClone(ctx, g, sid, v); err != nil {
		return sid, sv, err
	}
	return sid, sv, err
}

// ServiceID returns the ServiceID.
// It first checks the --service-id flag.
// Next it checks an environment variable.
// Next it checks the --service-name flag.
func ServiceID(ctx *cli.Context, g global.Data) (sid string, err error) {
	// FIXME: Implement manifest logic and lookup Service ID from manifest.
	if ctx.IsSet("service-id") {
		sid = ctx.Value("service-id").(string)
	} else if v := os.Getenv(env.ServiceID); v != "" {
		sid = v
	} else if ctx.IsSet("service-name") {
		paginator := g.APIClient.NewListServicesPaginator(&fastly.ListServicesInput{})
		var services []*fastly.Service
		for paginator.HasNext() {
			data, err := paginator.GetNext()
			if err != nil {
				return sid, fmt.Errorf("failed to identify Service ID from --service-name: %w", err)
			}
			services = append(services, data...)
		}
		var found bool
		for _, s := range services {
			if s.Name == ctx.Value("service-name").(string) {
				sid = s.ID
				found = true
				break
			}
		}
		if !found {
			return sid, errors.New("failed to match --service-name with an available service")
		}
	}

	return sid, err
}

// ServiceVersion returns the Service Version.
func ServiceVersion(ctx *cli.Context, g global.Data, sid string) (sv *fastly.Version, err error) {
	vs, err := g.APIClient.ListVersions(&fastly.ListVersionsInput{
		ServiceID: sid,
	})
	if err != nil || len(vs) == 0 {
		return sv, fmt.Errorf("failed to retrieve service version: %w", err)
	}

	// Sort versions into descending order.
	sort.Slice(vs, func(i, j int) bool {
		return vs[i].Number > vs[j].Number
	})

	var v *fastly.Version

	s := ctx.Value("service-version").(string)

	switch strings.ToLower(s) {
	case "latest":
		return vs[0], nil
	case "active":
		v, err = GetActiveVersion(vs)
	case "": // no --service-version flag provided
		v, err = GetActiveVersion(vs)
		if err != nil {
			return vs[0], nil // if no active version, return latest version
		}
	default:
		v, err = GetSpecifiedVersion(vs, s)
	}
	if err != nil {
		return sv, err
	}

	return v, nil
}

// GetActiveVersion returns the active service version.
func GetActiveVersion(vs []*fastly.Version) (*fastly.Version, error) {
	for _, v := range vs {
		if v.Active {
			return v, nil
		}
	}
	return nil, fmt.Errorf("failed to retrieve an active service version")
}

// GetSpecifiedVersion returns the specified service version.
func GetSpecifiedVersion(vs []*fastly.Version, version string) (*fastly.Version, error) {
	i, err := strconv.Atoi(version)
	if err != nil {
		return nil, err
	}

	for _, v := range vs {
		if v.Number == i {
			return v, nil
		}
	}

	return nil, fmt.Errorf("failed to retrieve specified service version: %s", version)
}

// AutoClone returns either the provided Service Version or an updated version.
// It will clone the service if the current version is either active or locked.
func AutoClone(ctx *cli.Context, g global.Data, sid string, v *fastly.Version) (sv int, err error) {
	sv = v.Number
	if ctx.IsSet("autoclone") {
		clonedVersion, err := g.APIClient.CloneVersion(&fastly.CloneVersionInput{
			ServiceID:      sid,
			ServiceVersion: sv,
		})
		if err != nil {
			return sv, fmt.Errorf("failed to clone service version '%d': %w", v.Number, err)
		}
		if ctx.IsSet("verbose") {
			fmt.Fprintf(ctx.Command.Writer, "Service version '%d' is not editable, so it was automatically cloned because --autoclone was enabled. Now operating on version '%d'.", v.Number, clonedVersion.Number)
		}
		return clonedVersion.Number, nil
	}
	// FIXME: opts.AllowActiveLocked
	if v.Active || v.Locked {
		return sv, fmt.Errorf("service version '%d' is not editable", sv)
	}
	return sv, err
}
