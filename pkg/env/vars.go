// Package env defines supported environment variables.
package env

const (
	// Token is the env var we look in for the Fastly API token.
	// G101 (CWE-798): Potential hardcoded credentials.
	// #nosec
	Token = "FASTLY_API_TOKEN"

	// Endpoint is the env var we look in for the API endpoint.
	Endpoint = "FASTLY_API_ENDPOINT"

	// ServiceID is the env var we look in for the required Service ID.
	ServiceID = "FASTLY_SERVICE_ID"

	// CustomerID is the env var we look in for a Customer ID.
	CustomerID = "FASTLY_CUSTOMER_ID"
)
