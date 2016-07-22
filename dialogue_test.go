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
	keys := []string{apiKey, ""}

	for i, key := range keys {
		d, err := NewDialogue(key)
		if err != nil {
			t.Logf("[%d] Expected error: %s", i, err.Error())
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

func TestNewDialogueWithInvalidHttpClient(t *testing.T) {

	apiKey := "valid-key"
	d, err := NewDialogue(apiKey, WithHttpClient(nil))
	if err != nil {
		t.Logf("Expected error: %s", err.Error())
	}
	if d != nil {
		t.Errorf("Expected nil, but got %v", d)
	}
}

func TestTalk(t *testing.T) {

	apiKey := os.Getenv("DOCOMO_API_KEY")

	d, err := NewDialogue(apiKey)
	if err != nil {
		t.Error(err)
	}
	r, err := d.Talk("今日の天気はどうですか？")
	if err != nil {
		t.Error(err)
	}
	t.Logf("response: %v", r)
}
