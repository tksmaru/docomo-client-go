package docomo

import (
	"fmt"
)

const (
	morphologicalEndpointForCorp       = "/gooLanguageAnalysisCorp/v1/morph"
	morphologicalEndpointForIndividual = "/gooLanguageAnalysis/v1/morph"
)

type MorphologicalRequest struct {
	RequestId  string `json:"request_id,omitempty"`
	Sentence   string `json:"sentence"`
	InfoFilter string `json:"info_filter,omitempty"`
	PosFilter  string `json:"pos_filter,omitempty"`
}

type MorphologicalResponse struct {
	RequestId  string       `json:"request_id"`
	InfoFilter string       `json:"info_filter,omitempty"`
	PosFilter  string       `json:"pos_filter,omitempty"`
	WordList   [][][]string `json:"word_list"`
}

type Morphological struct {
	client   *Client
	Endpoint string
}

func newMorphological(c *Client) *Morphological {
	m := &Morphological{
		client: c,
	}
	if c.settings.asCorp {
		m.Endpoint = fmt.Sprintf("%s%s?APIKEY=%s", apiDomain, morphologicalEndpointForCorp, c.APIKey)
	} else {
		m.Endpoint = fmt.Sprintf("%s%s?APIKEY=%s", apiDomain, morphologicalEndpointForIndividual, c.APIKey)
	}
	return m
}

func (m *Morphological) Post(req *MorphologicalRequest) (*MorphologicalResponse, error) {

	if req == nil {
		return nil, errInvalidRequest
	}
	var res *MorphologicalResponse
	if err := m.client.post(m.Endpoint, "application/json", req, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Morphological) Analyze(sentence string) (*MorphologicalResponse, error) {

	req := &MorphologicalRequest{
		Sentence: sentence,
	}
	return m.Post(req)
}
