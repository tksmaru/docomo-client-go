package docomo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	NamedEntityRequestUrlForCorp       = "https://api.apigw.smt.docomo.ne.jp/gooLanguageAnalysisCorp/v1/entity?APIKEY=%s"
	NamedEntityRequestUrlForIndividual = "https://api.apigw.smt.docomo.ne.jp/gooLanguageAnalysis/v1/entity?APIKEY=%s"
)

type NamedEntityRequest struct {
	RequestId   string `json:"request_id,omitempty"`
	Sentence    string `json:"sentence"`
	ClassFilter string `json:"class_filter,omitempty"`
}

type NamedEntityResponse struct {
	request_id  string     `json:"request_id`
	ClassFilter string     `json:"class_filter,omitempty"`
	NeList      [][]string `json:"ne_list"`
}

type NamedEntity struct {
	APIKey string
	*Settings
	Endpoint string
}

// for individual users
func NewNamedEntityForIndividual(APIKey string, options ...Option) (*NamedEntity, error) {
	return newNamedEntity(APIKey, NamedEntityRequestUrlForIndividual, options)
}

// for corporation users
func NewNamedEntityForCorp(APIKey string, options ...Option) (*NamedEntity, error) {
	return newNamedEntity(APIKey, NamedEntityRequestUrlForCorp, options)
}

func newNamedEntity(APIKey string, Endpoint string, options []Option) (*NamedEntity, error) {
	if !isValidKey(APIKey) {
		return nil, errInvalidApiKey
	}
	n := &NamedEntity{
		APIKey:   APIKey,
		Settings: NewSettings(),
		Endpoint: Endpoint,
	}
	if err := setOptions(n.Settings, options); err != nil {
		return nil, err
	}
	return n, nil
}

func (n *NamedEntity) Request(req *NamedEntityRequest) (*NamedEntityResponse, error) {

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response, err := n.client.Post(fmt.Sprintf(n.Endpoint, n.APIKey), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var res *NamedEntityResponse
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

func (n *NamedEntity) Extract(sentence string) (*NamedEntityResponse, error) {

	req := &NamedEntityRequest{
		Sentence: sentence,
	}
	return n.Request(req)
}
