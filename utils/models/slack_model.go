package models

type SlackMessageField struct {
	Title string      `json:"title"`
	Value interface{} `json:"value"`
}
