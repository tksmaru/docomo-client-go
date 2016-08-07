package docomo

import (
	"os"
	"testing"
)

func TestDialogue_Talk(t *testing.T) {

	apiKey := os.Getenv("DOCOMO_API_KEY")

	c, _ := NewClient(apiKey)
	r, err := c.Dialogue.Talk("今日の天気は")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("success response: %v", r)
}

func TestDialogue_Post_ErrorWithInvalidArgs(t *testing.T) {

	d, err := NewClient("__dummy__api__key__")
	if err != nil {
		t.Fatal(err)
	}
	r, err := d.Dialogue.Post(nil)
	if r != nil {
		t.Errorf("Expected nil, but got %v", r)
	}
	if err != errInvalidRequest {
		t.Errorf("Expected %s, but got %s", errInvalidRequest, err.Error())
	}

}

func TestDialogue_Talk_ErrorWithInvalidApiKey(t *testing.T) {

	d, err := NewClient("__invalid__api__key__")
	if err != nil {
		t.Fatal(err)
	}
	r, err := d.Dialogue.Talk("今日の天気はどうですか？")
	if r != nil {
		t.Errorf("Expected nil, but got %v", r)
	}
	expected := "POLSLA009: Unable to perform ApiKey based Authentication"
	if err.Error() != expected {
		t.Errorf("Expected %s, but got %s", expected, err.Error())
	}
}
