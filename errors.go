package docomo

import "errors"

var errInvalidApiKey = errors.New("Invalid API key.")
var errInvalidOption = errors.New("Invalid option.")
var errInvalidHttpClient = errors.New("Invalid http client.")
var errInvalidRequest = errors.New("Invalid request object.")

type CommonError struct {
	RequestError struct {
		PolicyException struct {
			MessageId string `json:"messageId"`
			Text      string `json:"text"`
		} `json:"policyException"`
	} `json:"requestError"`
}
