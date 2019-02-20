package dto

type Response struct {
	StatusCode   int         `json:"status_code"`
	Message      string      `json:"message"`
	DebugMessage string      `json:"debug_message,omitempty"`
	Data         interface{} `json:"data, omitempty"`
}
