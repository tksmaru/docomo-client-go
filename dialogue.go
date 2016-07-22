package docomo

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const dialogueURL = "https://api.apigw.smt.docomo.ne.jp/dialogue/v1/dialogue?APIKEY=%s"

type Dialogue struct {
	APIKey string
	*Settings
}

type DialogueRequest struct {
	Utt            string `json:"utt"`
	Context        string `json:"context,omitempty"`
	Nickname       string `json:"nickname,omitempty"`
	NicknameYomi   string `json:"nickname_y,omitempty"`
	Sex            string `json:"sex,omitempty"`
	BloodType      string `json:"bloodtype,omitempty"`
	BirthDateYear  int    `json:"birthdateY,string,omitempty"`
	BirthDateMonth int    `json:"birthdateM,string,omitempty"`
	BirthDateDay   int    `json:"birthdateD,string,omitempty"`
	Age            int    `json:"age,string,omitempty"`
	Constellations string `json:"constellations,string,omitempty"`
	Place          string `json:"place,string,omitempty"`
	Mode           string `json:"mode,string,omitempty"`
	Type           int    `json:"type,string,omitempty"`
}

type DialogueResponse struct {
	Utt     string `json:"utt"`
	Yomi    string `json:"yomi"`
	Mode    string `json:"mode"`
	Da      string `json:"da"`
	Context string `json:"context"`
}

// Initialize new dialogue instance
func NewDialogue(apiKey string, options ...Option) (*Dialogue, error) {

	if apiKey == "" {
		return nil, fmt.Errorf("Invalid API key: %v", apiKey)
	}

	d := &Dialogue{
		APIKey:   apiKey,
		Settings: NewSettings(),
	}

	for _, option := range options {
		if err := option(d.Settings); err != nil {
			return nil, err
		}
	}
	return d, nil
}

func (d *Dialogue) Request(req *DialogueRequest) (*DialogueResponse, error) {

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	fmt.Printf("request: %s", string(b)) // TODO delete

	response, err := d.client.Post(fmt.Sprintf(dialogueURL, d.APIKey), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res *DialogueResponse
	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func (d *Dialogue) Talk(phrase string) (*DialogueResponse, error) {

	req := &DialogueRequest{
		Utt: phrase,
	}
	return d.Request(req)
}
