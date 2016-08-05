package docomo

import (
	"os"
	"testing"
)

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
