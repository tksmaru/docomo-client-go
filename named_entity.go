package docomo

type NamedEntityRequest struct {
	RequestId   string `json:"request_id,omitempty"`
	Sentence    string `json:"sentence"`
	ClassFilter string `json:"class_filter,omitempty"`
}

type NamedEntityResponse struct {
	request_id  string `json:"request_id`
	ClassFilter string `json:"class_filter,omitempty"`
	NeList      string `json:"ne_list"` // TODO Not yet fixed
}

type NamedEntity struct {
	APIKey string
	Corporation bool // TODO Not yet fixed
	*Settings
}
