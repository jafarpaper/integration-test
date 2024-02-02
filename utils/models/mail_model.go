package models

type Mail struct {
	Template string                 `json:"template"`
	Data     map[string]interface{} `json:"data"`
	From     string                 `json:"from"`
	Name     string                 `json:"name"`
	To       string                 `json:"to"`
	Subject  string                 `json:"subject"`
}

type Whatsapp struct {
	Phone      string   `json:"phone"`
	Template   string   `json:"template"`
	Parameters []string `json:"parameters"`
}
