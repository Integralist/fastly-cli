// Package config manages global CLI configuration.
package config

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

// Filename is the name of the application configuration file.
const Filename = "config.toml"

// static is the embedded configuration file used as a failover.
//
//go:embed config.toml
var static []byte

// Path is the location of the application configuration file.
var Path = func() string {
	if dir, err := os.UserConfigDir(); err == nil {
		return filepath.Join(dir, "fastly", Filename)
	}
	if dir, err := os.UserHomeDir(); err == nil {
		return filepath.Join(dir, ".fastly", Filename)
	}
	panic("unable to deduce user config dir or user home dir")
}()

// Data represents our application toml configuration.
type Data struct {
	// ConfigVersion is the version of the configuration schema.
	ConfigVersion int `toml:"config_version"`
	// Fastly represents fastly specific configuration.
	Fastly Fastly `toml:"fastly"`
	// Language represents C@E language specific configuration.
	Language Language `toml:"language"`
	// Profiles represents multiple profile accounts.
	Profiles Profiles `toml:"profile"`
	// StarterKits represents language specific starter kits.
	StarterKits StarterKitLanguages `toml:"starter-kits"`
	// Viceroy represents viceroy specific configuration.
	Viceroy Viceroy `toml:"viceroy"`
}

// Read opens the file path and stores the content.
func (c *Data) Read(path string) error {
	f, err := os.Open(path)
	if err != nil {
		err = toml.Unmarshal(static, c)
		if err != nil {
			return fmt.Errorf("failed to decode internal static application config: %w", err)
		}
		return nil
	}
	defer f.Close()

	err = toml.NewDecoder(f).Decode(c)
	if err != nil {
		return fmt.Errorf("failed to decode internal application config: %w", err)
	}

	return nil
}

// CLI represents CLI specific configuration.
type CLI struct {
	Version string `toml:"version"`
}

// Fastly represents fastly specific configuration.
type Fastly struct {
	APIEndpoint string `toml:"api_endpoint"`
}

// Language represents C@E language specific configuration.
type Language struct {
	Go   Go   `toml:"go"`
	Rust Rust `toml:"rust"`
}

// Go represents Go C@E language specific configuration.
type Go struct {
	// TinyGoConstraint is the `tinygo` version that we support.
	TinyGoConstraint string `toml:"tinygo_constraint"`

	// ToolchainConstraint is the `go` version that we support.
	//
	// We aim for go versions that support go modules by default.
	// https://go.dev/blog/using-go-modules
	ToolchainConstraint string `toml:"toolchain_constraint"`
}

// Rust represents Rust C@E language specific configuration.
type Rust struct {
	// ToolchainConstraint is the `rustup` toolchain constraint for the compiler
	// that we support (a range is expected, e.g. >= 1.49.0 < 2.0.0).
	ToolchainConstraint string `toml:"toolchain_constraint"`

	// WasmWasiTarget is the Rust compilation target for Wasi capable Wasm.
	WasmWasiTarget string `toml:"wasm_wasi_target"`
}

// Profiles represents multiple profile accounts.
type Profiles map[string]*Profile

// Profile represents a specific profile account.
type Profile struct {
	Default bool   `toml:"default" json:"default"`
	Email   string `toml:"email" json:"email"`
	Token   string `toml:"token" json:"token"`
}

// StarterKitLanguages represents language specific starter kits.
type StarterKitLanguages struct {
	AssemblyScript []StarterKit `toml:"assemblyscript"`
	Go             []StarterKit `toml:"go"`
	JavaScript     []StarterKit `toml:"javascript"`
	Rust           []StarterKit `toml:"rust"`
}

// StarterKit represents starter kit specific configuration.
type StarterKit struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
	Path        string `toml:"path"`
	Tag         string `toml:"tag"`
	Branch      string `toml:"branch"`
}

// Viceroy represents viceroy specific configuration.
type Viceroy struct {
	// LastChecked is when the version of Viceroy was last checked.
	LastChecked string `toml:"last_checked"`
	// LatestVersion is the latest Viceroy version at the time it is set.
	LatestVersion string `toml:"latest_version"`
	// TTL is how long the CLI waits before considering the version stale.
	TTL string `toml:"ttl"`
}
