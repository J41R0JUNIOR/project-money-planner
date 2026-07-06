package aws

import (
	"os"
)

type Config struct {
	Region   string
	Profile  string
	Endpoint string
}

func LoadConfig() Config {
	return Config{
		Region:   firstNonEmpty(os.Getenv("AWS_REGION"), os.Getenv("AWS_DEFAULT_REGION"), "us-east-1"),
		Profile:  os.Getenv("AWS_PROFILE"),
		Endpoint: os.Getenv("AWS_ENDPOINT_URL"),
	}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
