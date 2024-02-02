package models

type PublisherMessage struct {
	Header  PublisherHeader `json:"header"`
	Content interface{}     `json:"content,omitempty"`
}

type PublisherHeader struct {
	RequestID  string      `json:"request_id"`
	Origin     string      `json:"origin"`
	ClientData interface{} `json:"client_data,omitempty"`
}
