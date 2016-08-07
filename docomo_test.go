package docomo

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {

	apiKey := "__dummy_valid_key__"
	c, err := NewClient(apiKey)
	if err != nil {
		t.Error(err)
	}
	expectedType := "*docomo.Client"
	if v := reflect.ValueOf(c).Type().String(); v != expectedType {
		t.Errorf("Expected %s, but got %s", expectedType, v)
	}
	if c.APIKey != apiKey {
		t.Errorf("Expected %s, but got %s", apiKey, c.APIKey)
	}
	if c.settings.asCorp {
		t.Errorf("Expected %t, but got %t", false, c.settings.asCorp)
	}

}

func TestNewClient_WithHttpClient(t *testing.T) {

	apiKey := "__dummy_valid_key__"
	timeout := time.Duration(3 * time.Second)
	hc := &http.Client{
		Timeout: timeout,
	}
	c, err := NewClient(apiKey, WithHttpClient(hc))
	if err != nil {
		t.Error(err)
	}
	expectedType := "*docomo.Client"
	if v := reflect.ValueOf(c).Type().String(); v != expectedType {
		t.Errorf("Expected %s, but got %s", expectedType, v)
	}
	if c.APIKey != apiKey {
		t.Errorf("Expected %s, but got %s", apiKey, c.APIKey)
	}
	if c.settings.client.Timeout != timeout {
		t.Errorf("Expected %d, but got %d", timeout, c.settings.client.Timeout)
	}
}

func TestNewClient_AsCorp(t *testing.T) {

	apiKey := "__dummy_valid_key__"
	c, err := NewClient(apiKey, AsCorp())
	if err != nil {
		t.Error(err)
	}
	expectedType := "*docomo.Client"
	if v := reflect.ValueOf(c).Type().String(); v != expectedType {
		t.Errorf("Expected %s, but got %s", expectedType, v)
	}
	if c.APIKey != apiKey {
		t.Errorf("Expected %s, but got %s", apiKey, c.APIKey)
	}
	if !c.settings.asCorp {
		t.Errorf("Expected %t, but got %t", true, c.settings.asCorp)
	}
}

func TestNewClient_WithHttpClientAndAsCorp(t *testing.T) {

	apiKey := "__dummy_valid_key__"
	timeout := time.Duration(3 * time.Second)
	hc := &http.Client{
		Timeout: timeout,
	}
	c, err := NewClient(apiKey, WithHttpClient(hc), AsCorp())
	if err != nil {
		t.Error(err)
	}
	expectedType := "*docomo.Client"
	if v := reflect.ValueOf(c).Type().String(); v != expectedType {
		t.Errorf("Expected %s, but got %s", expectedType, v)
	}
	if c.APIKey != apiKey {
		t.Errorf("Expected %s, but got %s", apiKey, c.APIKey)
	}
	if c.settings.client.Timeout != timeout {
		t.Errorf("Expected %d, but got %d", timeout, c.settings.client.Timeout)
	}
	if !c.settings.asCorp {
		t.Errorf("Expected %t, but got %t", true, c.settings.asCorp)
	}
}

func TestNewClient_Error_WithInvalidApiKey(t *testing.T) {

	var apiKey string
	invalidKeys := []string{apiKey, ""}

	for i, key := range invalidKeys {
		n, err := NewClient(key)
		if err == errInvalidApiKey {
			t.Logf("[%d] Expected error: %s", i, err.Error())
		} else if err != nil {
			t.Errorf("[%d] Expected %v, but got %v", i, errInvalidApiKey, err)
		}
		if n != nil {
			t.Errorf("[%d] Expected nil but got %v", i, n)
		}
	}
}

func TestNewClient_Error_WithInvalidOptions(t *testing.T) {

	invalidOptionsPattern := [][]Option{
		{nil},
		{nil, nil},
		{WithHttpClient(&http.Client{}), nil},
		{nil, WithHttpClient(&http.Client{})},
	}
	apiKey := "__dummy_valid_key__"

	for _, invalidOptions := range invalidOptionsPattern {
		n, err := NewClient(apiKey, invalidOptions...)
		if err == errInvalidOption {
			t.Logf("Expected error: %s", err.Error())
		} else if err != nil {
			t.Errorf("Expected %v, but got %v", errInvalidOption, err)
		}
		if n != nil {
			t.Errorf("Expected nil, but got %v", n)
		}
	}
}

func TestNewClient_Error_WithInvalidHttpClient(t *testing.T) {

	apiKey := "__dummy_valid_key__"
	n, err := NewClient(apiKey, WithHttpClient(nil))
	if err == errInvalidHttpClient {
		t.Logf("Expected error: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidHttpClient, err)
	}
	if n != nil {
		t.Errorf("Expected nil, but got %v", n)
	}
}
