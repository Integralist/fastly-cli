// Package global defines a globally accessible data structure.
package global

import (
	"log"
	"os"

	"github.com/fastly/go-fastly/v8/fastly"

	"github.com/integralist/fastly-cli/pkg/config"
)

// Data is shared globally across the application.
type Data struct {
	// APIClient is used to interact with the Fastly API.
	APIClient *fastly.Client
	// Config is the application configuration.
	Config *config.Data
}

// Container represents a single instance of our global data.
var Container Data

func init() {
	client, err := fastly.NewClient(os.Getenv("FASTLY_API_TOKEN"))
	if err != nil {
		log.Fatalf("failed to instantiate Fastly API client: %s", err)
	}

	var cfg config.Data
	err = cfg.Read(config.Path)
	if err != nil {
		log.Fatalf("failed to read application configuration: %s", err)
	}

	Container = Data{
		APIClient: client,
		Config:    &cfg,
	}
}
