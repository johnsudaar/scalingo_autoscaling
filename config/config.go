package config

import (
	"crypto/tls"

	"github.com/Scalingo/go-scalingo"
)

var (
	API_TOKEN        = "nojJkhM8gkWmt_eWMCP6HY4-A9p9Q8e9Oh3OQFfTX8Y"
	SCALINGO_API_URL = "https://api.scalingo.com"
	MAX_CONTAINER    = 3
	MIN_CONTAINER    = 1

	// NEURAL
	NEURAL_EPSILON = float64(1)
	NEURAL_HIGH    = float64(0.8)
	NEURAL_LOW     = float64(0.5)
)

func ScalingoClient() *scalingo.Client {
	return scalingo.NewClient(scalingo.ClientConfig{
		APIToken:  API_TOKEN,
		Endpoint:  SCALINGO_API_URL,
		TLSConfig: &tls.Config{},
	})
}
