package docomo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//const dialogueEndpoint = "https://api.apigw.smt.docomo.ne.jp/dialogue/v1/dialogue?APIKEY=%s"
const dialogueEndpoint = "/dialogue/v1/dialogue"

type Dialogue struct {
	//APIKey string
	//*Settings
	*Client
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
func newDialogue( /*apiKey string, options ...Option*/ client *Client) *Dialogue /*, error*/ {

	//if !isValidKey(apiKey) {
	//	return nil, errInvalidApiKey
	//}

	d := &Dialogue{
		//APIKey:   apiKey,
		//Settings: NewSettings(),
		client:   client,
		Endpoint: fmt.Sprintf("%s%s?APIKEY=%s", apiDomain, dialogueEndpoint, client.APIKey),
	}
	//if err := setOptions(d.Settings, options); err != nil {
	//	return nil, err
	//}
	//return d, nil
	return d
}

func (d *Dialogue) Request(req *DialogueRequest) (*DialogueResponse, error) {

	if req == nil {
		return nil, errInvalidRequest
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response, err := d.post(d.Endpoint, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var res *DialogueResponse
		if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
			return nil, err
		}
		return res, nil
	} else {
		var errorRes *CommonError
		if err := json.NewDecoder(response.Body).Decode(&errorRes); err != nil {
			return nil, err
		}
		e := errorRes.RequestError.PolicyException
		return nil, fmt.Errorf("%s: %s", e.MessageId, e.Text)
	}
}

func (d *Dialogue) Talk(phrase string) (*DialogueResponse, error) {

	req := &DialogueRequest{
		Utt: phrase,
	}
	return d.Request(req)
}
