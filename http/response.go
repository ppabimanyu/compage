package http

type Response struct {
	StatusCode int    `json:"status_code" xml:"status_code"`
	RequestID  string `json:"request_id" xml:"request_id"`
	Message    string `json:"message" xml:"message"`
	Error      *Error `json:"error" xml:"error"`
	Data       any    `json:"data" xml:"data"`
}

type Error struct {
	Code    string `json:"error_code" xml:"error_code"`
	Details any    `json:"details" xml:"details"`
}
