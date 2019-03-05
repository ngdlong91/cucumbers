package dto

var DebugResponse bool

func SetDebugResponse(val bool) {
	DebugResponse = val
}

type Response struct {
	StatusCode   int         `json:"status_code"`
	Message      string      `json:"message"`
	DebugMessage string      `json:"debug_message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func (r *Response) SetDebugMsg(msg string) *Response {
	if DebugResponse {
		r.DebugMessage = msg
	}
	return r
}
