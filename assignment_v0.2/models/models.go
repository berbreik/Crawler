package models

// ResponseData is the response data
type ResponseData struct {
	Result []Body `json:"result"`
	Error  error `json:"error"`
}

// RequestData is the request data
type RequestData struct {
	Urls []string `json:"urls"`
}

// Body is the body of the response
type Body struct {
	Url  string `json:"url"`
	Data string `json:"data"`
}
