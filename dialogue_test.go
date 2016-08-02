package docomo

import (
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestNewDialogue(t *testing.T) {

	apiKey := "valid-key"
	d, err := NewDialogue(apiKey)
	if err != nil {
		t.Error(err)
	}
	expectedType := "*docomo.Dialogue"
	if v := reflect.ValueOf(d).Type().String(); v != expectedType {
		t.Errorf("Expected %s, but got %s", expectedType, v)
	}
	if d.APIKey != apiKey {
		t.Errorf("Expected %s, but got %s", apiKey, d.APIKey)
	}
}

func TestNewDialogueWithInvalidAPIKey(t *testing.T) {

	var apiKey string
	invalidKeys := []string{apiKey, ""}

	for i, key := range invalidKeys {
		d, err := NewDialogue(key)
		if err == errInvalidApiKey {
			t.Logf("[%d] Expected error: %s", i, err.Error())
		} else if err != nil {
			t.Errorf("[%d] Expected %v, but got %v", i, errInvalidApiKey, err)
		}
		if d != nil {
			t.Errorf("[%d] Expected nil but got %v", i, d)
		}
	}
}

func TestNewDialogueWithHttpClient(t *testing.T) {

	apiKey := "valid-key"
	timeout := time.Duration(3 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	d, err := NewDialogue(apiKey, WithHttpClient(client))
	if err != nil {
		t.Error(err)
	}
	expectedType := "*docomo.Dialogue"
	if v := reflect.ValueOf(d).Type().String(); v != expectedType {
		t.Errorf("Expected %s, but got %s", expectedType, v)
	}
	if d.APIKey != apiKey {
		t.Errorf("Expected %s, but got %s", apiKey, d.APIKey)
	}
	if d.client.Timeout != timeout {
		t.Errorf("Expected %d, but got %d", timeout, d.client.Timeout)
	}
}

func TestNewDialogueWithInvalidOptions(t *testing.T) {

	invalidOptionsPattern := [][]Option{
		{nil},
		{nil, nil},
		{WithHttpClient(&http.Client{}), nil},
		{nil, WithHttpClient(&http.Client{})},
	}
	apiKey := "valid-key"

	for _, invalidOptions := range invalidOptionsPattern {
		d, err := NewDialogue(apiKey, invalidOptions...)
		if err == errInvalidOption {
			t.Logf("Expected error: %s", err.Error())
		} else if err != nil {
			t.Errorf("Expected %v, but got %v", errInvalidOption, err)
		}
		if d != nil {
			t.Errorf("Expected nil, but got %v", d)
		}
	}
}

func TestNewDialogueWithInvalidHttpClient(t *testing.T) {

	apiKey := "valid-key"
	d, err := NewDialogue(apiKey, WithHttpClient(nil))
	if err == errInvalidHttpClient {
		t.Logf("Expected error: %s", err.Error())
	} else if err != nil {
		t.Errorf("Expected %v, but got %v", errInvalidHttpClient, err)
	}
	if d != nil {
		t.Errorf("Expected nil, but got %v", d)
	}
}

func TestTalk(t *testing.T) {

	apiKey := os.Getenv("DOCOMO_DIALOGUE_API_KEY")

	d, err := NewDialogue(apiKey)
	if err != nil {
		t.Fatal(err)
	}
	r, err := d.Talk("今日の天気はどうですか？")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("success response: %v", r)
}

func TestTalkErrorWithInvalidApiKey(t *testing.T) {

	d, err := NewDialogue("__invalid__api__key__")
	if err != nil {
		t.Fatal(err)
	}
	r, err := d.Talk("今日の天気はどうですか？")
	if r != nil {
		t.Errorf("Expected nil, but got %v", r)
	}
	expected := "POLSLA009: Unable to perform ApiKey based Authentication"
	if err.Error() != expected {
		t.Errorf("Expected %s, but got %s", expected, err.Error())
	}
}

func TestRequestErrorWithInvalidRequest(t *testing.T) {

	d, err := NewDialogue("__dummy__api__key__")
	if err != nil {
		t.Fatal(err)
	}
	r, err := d.Request(nil)
	if r != nil {
		t.Errorf("Expected nil, but got %v", r)
	}
	if err != errInvalidDialogueRequest {
		t.Errorf("Expected %v, but got %v", errInvalidDialogueRequest, err)
	}
}
