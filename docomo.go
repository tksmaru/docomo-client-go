package docomo

import (
	"fmt"
	"net/http"
)

type Settings struct {
	client *http.Client
}

func NewSettings() *Settings {
	return &Settings{
		client: http.DefaultClient,
	}
}

// Optional settings.
type Option func(s *Settings) error

// Configure http client
func WithHttpClient(client *http.Client) Option {
	return func(s *Settings) error {
		if client == nil {
			return fmt.Errorf("Invalid http client: nil")
		}
		s.client = client
		return nil
	}
}
