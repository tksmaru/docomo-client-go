package docomo

import (
	"testing"
	"os"
)

func TestDialogue(t *testing.T) {

	APIKey := os.Getenv("DOCOMO_API_KEY")

	d, err := NewDocomo(APIKey)
	if err != nil {
		t.Error(err)
	}
	r, err := d.Dialogue("今日の天気はどうですか？")
	if err != nil {
		t.Error(err)
	}
	t.Logf("response: %v", r)
}
