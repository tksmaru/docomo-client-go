package docomo

import (
	"errors"
	"io"
	"net/http"
)

var errInvalidApiKey = errors.New("Invalid API key.")
var errInvalidOption = errors.New("Invalid option.")
var errInvalidHttpClient = errors.New("Invalid http client.")
var errInvalidRequest = errors.New("Invalid request object.")

const apiDomain = "https://api.apigw.smt.docomo.ne.jp"

type Client struct {
	APIKey   string
	settings *Settings
	Dialogue *Dialogue
}

func NewClient(apiKey string, options ...Option) (*Client, error) {

	if !isValidKey(apiKey) {
		return nil, errInvalidApiKey
	}

	c := &Client{
		APIKey:   apiKey,
		settings: NewSettings(),
		//Dialogue: NewDialogue(),
	}
	if err := setOptions(c.settings, options); err != nil {
		return nil, err
	}

	c.Dialogue = newDialogue(c)

	return c, nil
}

type Settings struct {
	client *http.Client
	isCorp bool
}

// Default Settings for client.
func NewSettings() *Settings {
	return &Settings{
		client: http.DefaultClient,
		isCorp: false,
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

// use as a corporation account
func AsCorp() Option {
	return func(s *Settings) error {
		s.isCorp = true
		return nil
	}
}

// Validate Keys. This validation checks for nil or empty string.
func isValidKey(apiKey string) bool {
	if apiKey == "" {
		return false
	}
	return true
}

func (c *Client) post(url string, bodyType string, body io.Reader) (*http.Response, error) {
	return c.settings.client.Post(url, bodyType, body)
}
