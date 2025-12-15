package challenge

import "net/http"

type HTTPRequest struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    []byte
}

type Option func(*HTTPRequest)

func WithMethod(m string) Option {
	return func(h *HTTPRequest) {
		h.Method = m
	}
}

func WithHeader(key, value string) Option {
	return func(h *HTTPRequest) {
		h.Headers[key] = value
	}
}

func WithBody(jsonStr string) Option {
	return func(h *HTTPRequest) {
		h.Body = []byte(jsonStr)
	}
}

func NewRequest(url string, opts ...Option) *HTTPRequest {
	request := &HTTPRequest{
		URL:     url,
		Method:  http.MethodGet,
		Headers: make(map[string]string),
		Body:    nil,
	}

	for _, opt := range opts {
		opt(request)
	}

	return request
}
