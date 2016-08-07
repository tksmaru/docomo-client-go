package docomo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const apiDomain = "https://api.apigw.smt.docomo.ne.jp"

type Client struct {
	APIKey        string
	settings      *Settings
	Dialogue      *Dialogue
	NamedEntity   *NamedEntity
	Morphological *Morphological
}

func NewClient(apiKey string, options ...Option) (*Client, error) {

	if !isValidKey(apiKey) {
		return nil, errInvalidApiKey
	}

	c := &Client{
		APIKey:   apiKey,
		settings: NewSettings(),
	}
	if err := setOptions(c.settings, options); err != nil {
		return nil, err
	}

	c.Dialogue = newDialogue(c)
	c.NamedEntity = newNamedEntity(c)
	c.Morphological = newMorphological(c)

	return c, nil
}

type Settings struct {
	client *http.Client
	asCorp bool
}

// Default Settings for client.
func NewSettings() *Settings {
	return &Settings{
		client: http.DefaultClient,
		asCorp: false,
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
		s.asCorp = true
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

func (c *Client) post(url string, bodyType string, req interface{}, res interface{}) error {

	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := c.settings.client.Post(url, bodyType, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
			return err
		}
		return nil
	} else {
		var errorRes *CommonError
		if err := json.NewDecoder(resp.Body).Decode(&errorRes); err != nil {
			return err
		}
		e := errorRes.RequestError.PolicyException
		return fmt.Errorf("%s: %s", e.MessageId, e.Text)
	}
}
