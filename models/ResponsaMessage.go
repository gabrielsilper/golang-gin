package models

type ResponseMessage struct {
	Message string `json:"message"`
}

func NewResponseMessage(message string) ResponseMessage {
	return ResponseMessage{message}
}
