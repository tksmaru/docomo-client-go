package docomo

type PolicyException struct {
	MessageId string `json:"messageId"`
	Text      string `json:"text"`
}

type RequestError struct {
	PolicyException PolicyException `json:"policyException"`
}

type CommonError struct {
	RequestError RequestError `json:"requestError"`
}

