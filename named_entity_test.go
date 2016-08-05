package docomo

import (
	"os"
	"testing"
	"reflect"
	"net/http"
	"time"
)

func TestNewNamedEntityForIndividual(t *testing.T) {

	apiKey := "valid-key"
	n, err := NewNamedEntityForIndividual(apiKey)
	if err != nil {
		t.Error(err)
	}
	expectedType := "*docomo.NamedEntity"
	if v := reflect.ValueOf(n).Type().String(); v != expectedType {
		t.Errorf("Expected %s, but got %s", expectedType, v)
	}
	if n.APIKey != apiKey {
		t.Errorf("Expected %s, but got %s", apiKey, n.APIKey)
	}
	if n.Endpoint != namedEntityEndpointForIndividual {
		t.Errorf("Expected %s, but got %s", namedEntityEndpointForIndividual, n.Endpoint)
	}
}

func TestNewNamedEntityForCorp(t *testing.T) {

	apiKey := "valid-key"
	n, err := NewNamedEntityForCorp(apiKey)
	if err != nil {
		t.Error(err)
	}
	expectedType := "*docomo.NamedEntity"
	if v := reflect.ValueOf(n).Type().String(); v != expectedType {
		t.Errorf("Expected %s, but got %s", expectedType, v)
	}
	if n.APIKey != apiKey {
		t.Errorf("Expected %s, but got %s", apiKey, n.APIKey)
	}
	if n.Endpoint != namedEntityEndpointForCorp {
		t.Errorf("Expected %s, but got %s", namedEntityEndpointForCorp, n.Endpoint)
	}
}

func TestNewNamedEntityWithHttpClient(t *testing.T) {

	apiKey := "valid-key"
	timeout := time.Duration(3 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	n, err := NewNamedEntityForIndividual(apiKey, WithHttpClient(client))
	if err != nil {
		t.Error(err)
	}
	expectedType := "*docomo.NamedEntity"
	if v := reflect.ValueOf(n).Type().String(); v != expectedType {
		t.Errorf("Expected %s, but got %s", expectedType, v)
	}
	if n.APIKey != apiKey {
		t.Errorf("Expected %s, but got %s", apiKey, n.APIKey)
	}
	if n.client.Timeout != timeout {
		t.Errorf("Expected %d, but got %d", timeout, n.client.Timeout)
	}
}

func TestNewNamedEntityWithInvalidAPIKey(t *testing.T) {

	var apiKey string
	invalidKeys := []string{apiKey, ""}

	for i, key := range invalidKeys {
		n, err := NewNamedEntityForIndividual(key)
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

func TestNewNamedEntityWithInvalidOptions(t *testing.T) {

	invalidOptionsPattern := [][]Option{
		{nil},
		{nil, nil},
		{WithHttpClient(&http.Client{}), nil},
		{nil, WithHttpClient(&http.Client{})},
	}
	apiKey := "valid-key"

	for _, invalidOptions := range invalidOptionsPattern {
		n, err := NewNamedEntityForIndividual(apiKey, invalidOptions...)
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

func TestNewNamedEntityWithInvalidHttpClient(t *testing.T) {

	apiKey := "valid-key"
	n, err := NewNamedEntityForIndividual(apiKey, WithHttpClient(nil))
	if err == errInvalidHttpClient {
		t.Logf("Expected error: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidHttpClient, err)
	}
	if n != nil {
		t.Errorf("Expected nil, but got %v", n)
	}
}

func TestExtract(t *testing.T) {

	apiKey := os.Getenv("DOCOMO_API_KEY")

	n, err := NewNamedEntityForIndividual(apiKey)
	if err != nil {
		t.Fatal(err)
	}
	r, err := n.Extract("今日の5時の千葉の天気を千葉県庁の佐藤さんが確認した")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("success response: %v", r)
}

func TestExtractErrorWithInvalidApiKey(t *testing.T) {

	n, err := NewNamedEntityForIndividual("__invalid__api__key__")
	if err != nil {
		t.Fatal(err)
	}
	r, err := n.Extract("今日の天気はどうですか？")
	if r != nil {
		t.Errorf("Expected nil, but got %v", r)
	}
	expected := "POLSLA009: Unable to perform ApiKey based Authentication"
	if err.Error() != expected {
		t.Errorf("Expected %s, but got %s", expected, err.Error())
	}
}

func TestExtractErrorWithInvalidRequest(t *testing.T) {

	n, err := NewNamedEntityForIndividual("__dummy__api__key__")
	if err != nil {
		t.Fatal(err)
	}
	r, err := n.Request(nil)
	if r != nil {
		t.Errorf("Expected nil, but got %v", r)
	}
	if err != errInvalidRequest {
		t.Errorf("Expected %v, but got %v", errInvalidRequest, err)
	}
}
