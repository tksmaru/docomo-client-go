package docomo

import (
	"testing"
	"os"
)

func TestTalk(t *testing.T) {

	APIKey := os.Getenv("DOCOMO_API_KEY")

	d, err := NewDialogue(APIKey)
	if err != nil {
		t.Error(err)
	}
	r, err := d.Talk("今日の天気はどうですか？")
	if err != nil {
		t.Error(err)
	}
	t.Logf("response: %v", r)
}
