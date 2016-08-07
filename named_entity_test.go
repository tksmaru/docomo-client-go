package docomo

import (
	"fmt"
	"os"
	"testing"
)

func TestNamedEntity_AsIndividual(t *testing.T) {

	apiKey := "__dummy__api__key__"
	c, _ := NewClient(apiKey)
	expected := fmt.Sprintf("%s%s?APIKEY=%s", apiDomain, namedEntityEndpointForIndividual, apiKey)
	if c.NamedEntity.Endpoint != expected {
		t.Errorf("Expected %s, but got %s", expected, c.NamedEntity.Endpoint)
	}
}

func TestNamedEntity_AsCorp(t *testing.T) {

	apiKey := "__dummy__api__key__"
	c, _ := NewClient(apiKey, AsCorp())
	expected := fmt.Sprintf("%s%s?APIKEY=%s", apiDomain, namedEntityEndpointForCorp, apiKey)
	if c.NamedEntity.Endpoint != expected {
		t.Errorf("Expected %s, but got %s", expected, c.NamedEntity.Endpoint)
	}
}

func TestNamedEntity_Extract(t *testing.T) {

	apiKey := os.Getenv("DOCOMO_API_KEY")

	c, _ := NewClient(apiKey)
	r, err := c.NamedEntity.Extract("今日の5時の千葉の天気を千葉県庁の佐藤さんが確認した")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("success response: %v", r)
}

func TestNamedEntity_Post_ErrorWithInvalidArgs(t *testing.T) {

	c, err := NewClient("__dummy__api__key__")
	if err != nil {
		t.Fatal(err)
	}
	r, err := c.NamedEntity.Post(nil)
	if r != nil {
		t.Errorf("Expected nil, but got %v", r)
	}
	if err != errInvalidRequest {
		t.Errorf("Expected %s, but got %s", errInvalidRequest, err.Error())
	}
}

func TestNamedEntity_Post_ErrorWithInvalidApiKey(t *testing.T) {

	c, err := NewClient("__invalid__api__key__")
	if err != nil {
		t.Fatal(err)
	}
	r, err := c.NamedEntity.Extract("今日の5時の千葉の天気を千葉県庁の佐藤さんが確認した")
	if r != nil {
		t.Errorf("Expected nil, but got %v", r)
	}
	expected := "POLSLA009: Unable to perform ApiKey based Authentication"
	if err.Error() != expected {
		t.Errorf("Expected %s, but got %s", expected, err.Error())
	}
}
