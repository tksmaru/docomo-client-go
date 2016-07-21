package docomo

import (
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
)

const dialogueURL = "https://api.apigw.smt.docomo.ne.jp/dialogue/v1/dialogue?APIKEY=%s"

type Docomo struct {
	APIKey string
	Client *http.Client
}


type DialogueRequest struct {
	Utt string `json:"utt"`
	Context string `json:"context,omitempty"`
	// TODO full params
}

type DialogueResponse struct {
	Utt string `json:"utt"`
	Yomi string `json:"yomi"`
	Mode string `json:"mode"`
	Da string `json:"da"`
	Context string `json:"context"`
}

func NewDocomo(APIKey string) (*Docomo, error) {
	d := &Docomo{
		APIKey: APIKey,
		Client: http.DefaultClient,
	}
	return d, nil
}

func (d *Docomo) DialogueFull(req *DialogueRequest) (*DialogueResponse, error) {

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	fmt.Printf("request: %s", string(b)) // TODO delete

	response, err := d.Client.Post(fmt.Sprintf(dialogueURL, d.APIKey), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res *DialogueResponse
	if err :=json.NewDecoder(response.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (d *Docomo) Dialogue(phrase string) (*DialogueResponse, error) {

	req := &DialogueRequest{
		Utt: phrase,
	}
	return d.DialogueFull(req)
}
