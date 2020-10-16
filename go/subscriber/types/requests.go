package types

type MessageRequest struct {
	Message  interface{} `json:"message" binding:"required"`
}
