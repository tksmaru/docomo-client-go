package docomo

import (
	"fmt"
)

const dialogueEndpoint = "/dialogue/v1/dialogue"

type Dialogue struct {
	client   *Client
	Endpoint string
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
func newDialogue(c *Client) *Dialogue {

	d := &Dialogue{
		client:   c,
		Endpoint: fmt.Sprintf("%s%s?APIKEY=%s", apiDomain, dialogueEndpoint, c.APIKey),
	}
	return d
}

func (d *Dialogue) Post(req *DialogueRequest) (*DialogueResponse, error) {

	if req == nil {
		return nil, errInvalidRequest
	}
	var res *DialogueResponse
	if err := d.client.post(d.Endpoint, "application/json", req, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (d *Dialogue) Talk(phrase string) (*DialogueResponse, error) {

	req := &DialogueRequest{
		Utt: phrase,
	}
	return d.Post(req)
}
