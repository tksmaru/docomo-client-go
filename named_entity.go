package docomo

import (
	"fmt"
)

const (
	namedEntityEndpointForCorp       = "/gooLanguageAnalysisCorp/v1/entity"
	namedEntityEndpointForIndividual = "/gooLanguageAnalysis/v1/entity"
)

type NamedEntityRequest struct {
	RequestId   string `json:"request_id,omitempty"`
	Sentence    string `json:"sentence"`
	ClassFilter string `json:"class_filter,omitempty"`
}

type NamedEntityResponse struct {
	RequestId   string     `json:"request_id"`
	ClassFilter string     `json:"class_filter,omitempty"`
	NeList      [][]string `json:"ne_list"`
}

type NamedEntity struct {
	client   *Client
	Endpoint string
}

func newNamedEntity(c *Client) *NamedEntity {
	n := &NamedEntity{
		client: c,
	}
	if c.settings.asCorp {
		n.Endpoint = fmt.Sprintf("%s%s?APIKEY=%s", apiDomain, namedEntityEndpointForCorp, c.APIKey)
	} else {
		n.Endpoint = fmt.Sprintf("%s%s?APIKEY=%s", apiDomain, namedEntityEndpointForIndividual, c.APIKey)
	}
	return n
}

func (n *NamedEntity) Post(req *NamedEntityRequest) (*NamedEntityResponse, error) {

	if req == nil {
		return nil, errInvalidRequest
	}
	var res *NamedEntityResponse
	if err := n.client.post(n.Endpoint, "application/json", req, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (n *NamedEntity) Extract(sentence string) (*NamedEntityResponse, error) {

	req := &NamedEntityRequest{
		Sentence: sentence,
	}
	return n.Post(req)
}
