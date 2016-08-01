package docomo

import (
	"errors"
	"net/http"
)

var errInvalidApiKey = errors.New("Invalid API key.")
var errInvalidOption = errors.New("Invalid option.")
var errInvalidHttpClient = errors.New("Invalid http client.")

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

// Set Optional settings.
func setOptions(s *Settings, options []Option) error {
	for _, option := range options {
		if option == nil {
			return errInvalidOption
		}
		if err := option(s); err != nil {
			return err
		}
	}
	return nil
}

// Configure http client
func WithHttpClient(client *http.Client) Option {
	return func(s *Settings) error {
		if client == nil {
			return errInvalidHttpClient
		}
		s.client = client
		return nil
	}
}
