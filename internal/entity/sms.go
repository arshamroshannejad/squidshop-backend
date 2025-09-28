package entity

type SmsRequest struct {
	Recipient []string `json:"recipient"`
	Sender    string   `json:"sender"`
	Message   string   `json:"message"`
}
