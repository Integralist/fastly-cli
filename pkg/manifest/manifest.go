package manifest

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Filename is the name of the package manifest file.
const Filename = "fastly.toml"

// File is the Compute package manifest data.
type File struct {
	// Authors is a list of project authors (typically an email).
	Authors []string `toml:"authors"`
	// Description is the project description.
	Description string `toml:"description"`
	// Language is the programming language used for the project.
	Language string `toml:"language"`
	// Profile is the name of the profile account the Fastly CLI should use to make API requests.
	Profile string `toml:"profile,omitempty"`
	// LocalServer describes the configuration for the local server built into the Fastly CLI.
	LocalServer LocalServer `toml:"local_server,omitempty"`
	// ManifestVersion is the manifest schema version number.
	ManifestVersion int `toml:"manifest_version"`
	// Name is the package name.
	Name string `toml:"name"`
	// Scripts describes customisation options for the Fastly CLI build step.
	Scripts Scripts `toml:"scripts,omitempty"`
	// ServiceID is the Fastly Service ID to deploy the package to.
	ServiceID string `toml:"service_id"`
	// Setup describes a set of service configuration that works with the code in the package.
	Setup Setup `toml:"setup,omitempty"`
}

// LocalServer represents a list of mocked Viceroy resources.
type LocalServer struct {
	Backends       map[string]LocalBackend       `toml:"backends"`
	ConfigStores   map[string]LocalConfigStore   `toml:"config_stores,omitempty"`
	KVStores       map[string][]LocalKVStore     `toml:"kv_stores,omitempty"`
	SecretStores   map[string][]LocalSecretStore `toml:"secret_stores,omitempty"`
	ViceroyVersion string                        `toml:"viceroy_version,omitempty"`
}

// LocalBackend represents a backend to be mocked by the local testing server.
type LocalBackend struct {
	URL          string `toml:"url"`
	OverrideHost string `toml:"override_host,omitempty"`
	CertHost     string `toml:"cert_host,omitempty"`
	UseSNI       bool   `toml:"use_sni,omitempty"`
}

// LocalConfigStore represents a config store to be mocked by the local testing server.
type LocalConfigStore struct {
	File     string            `toml:"file,omitempty"`
	Format   string            `toml:"format"`
	Contents map[string]string `toml:"contents,omitempty"`
}

// LocalKVStore represents an kv_store to be mocked by the local testing server.
type LocalKVStore struct {
	Key  string `toml:"key"`
	File string `toml:"file,omitempty"`
	Data string `toml:"data,omitempty"`
}

// LocalSecretStore represents a secret_store to be mocked by the local testing server.
type LocalSecretStore struct {
	Key  string `toml:"key"`
	File string `toml:"file,omitempty"`
	Data string `toml:"data,omitempty"`
}

// Scripts represents build configuration.
type Scripts struct {
	Build     string `toml:"build,omitempty"`
	PostBuild string `toml:"post_build,omitempty"`
	PostInit  string `toml:"post_init,omitempty"`
}

// Setup represents a set of service configuration that works with the code in
// the package. See https://developer.fastly.com/reference/fastly-toml/.
type Setup struct {
	Backends     map[string]*SetupBackend     `toml:"backends,omitempty"`
	ConfigStores map[string]*SetupConfigStore `toml:"config_stores,omitempty"`
	Loggers      map[string]*SetupLogger      `toml:"log_endpoints,omitempty"`
	ObjectStores map[string]*SetupKVStore     `toml:"object_stores,omitempty"`
	KVStores     map[string]*SetupKVStore     `toml:"kv_stores,omitempty"`
	SecretStores map[string]*SetupSecretStore `toml:"secret_stores,omitempty"`
}

// Defined indicates if there is any [setup] configuration in the manifest.
func (s Setup) Defined() bool {
	var defined bool

	if len(s.Backends) > 0 {
		defined = true
	}
	if len(s.ConfigStores) > 0 {
		defined = true
	}
	if len(s.Loggers) > 0 {
		defined = true
	}
	if len(s.KVStores) > 0 {
		defined = true
	}

	return defined
}

// SetupBackend represents a '[setup.backends.<T>]' instance.
type SetupBackend struct {
	Address     string `toml:"address,omitempty"`
	Port        int    `toml:"port,omitempty"`
	Description string `toml:"description,omitempty"`
}

// SetupConfigStore represents a '[setup.dictionaries.<T>]' instance.
type SetupConfigStore struct {
	Items       map[string]SetupConfigStoreItems `toml:"items,omitempty"`
	Description string                           `toml:"description,omitempty"`
}

// SetupConfigStoreItems represents a '[setup.dictionaries.<T>.items]' instance.
type SetupConfigStoreItems struct {
	Value       string `toml:"value,omitempty"`
	Description string `toml:"description,omitempty"`
}

// SetupLogger represents a '[setup.log_endpoints.<T>]' instance.
type SetupLogger struct {
	Provider string `toml:"provider,omitempty"`
}

// SetupKVStore represents a '[setup.kv_stores.<T>]' instance.
type SetupKVStore struct {
	Items       map[string]SetupKVStoreItems `toml:"items,omitempty"`
	Description string                       `toml:"description,omitempty"`
}

// SetupKVStoreItems represents a '[setup.kv_stores.<T>.items]' instance.
type SetupKVStoreItems struct {
	Value       string `toml:"value,omitempty"`
	Description string `toml:"description,omitempty"`
}

// SetupSecretStore represents a '[setup.secret_stores.<T>]' instance.
type SetupSecretStore struct {
	Entries     map[string]SetupSecretStoreEntry `toml:"entries,omitempty"`
	Description string                           `toml:"description,omitempty"`
}

// SetupSecretStoreEntry represents a '[setup.secret_stores.<T>.entries]' instance.
type SetupSecretStoreEntry struct {
	// The secret value is intentionally omitted to avoid secrets
	// from being included in the manifest. Instead, secret
	// values are input during setup.
	Description string `toml:"description,omitempty"`
}

// Read opens the file path and returns the content.
func Read(path string) (*File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open manifest: %w", err)
	}
	defer f.Close()

	var m File
	err = toml.NewDecoder(f).Decode(&m)
	if err != nil {
		return nil, fmt.Errorf("failed to decode manifest: %w", err)
	}

	return &m, nil
}
