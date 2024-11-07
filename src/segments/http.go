package segments

import (
	"encoding/json"

	"net/http"

	"github.com/jandedobbeleer/oh-my-posh/src/properties"
)

type HTTP struct {
	base

	Body map[string]interface{}
}

const (
	METHOD properties.Property = "method"
)

func (h *HTTP) Template() string {
	return " {{ .Body }} "
}

func (h *HTTP) Enabled() bool {
	url := h.props.GetString(URL, "")
	if len(url) == 0 {
		return false
	}

	method := h.props.GetString(METHOD, "GET")

	result, err := h.getResult(url, method)
	if err != nil {
		return false
	}

	h.Body = result
	return true
}

func (h *HTTP) getResult(url, method string) (map[string]interface{}, error) {
	setMethod := func(request *http.Request) {
		request.Method = method
	}

	resultBody, err := h.env.HTTPRequest(url, nil, 10000, setMethod)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(resultBody, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
