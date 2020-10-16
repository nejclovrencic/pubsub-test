package types

type SubscribeRequest struct {
	Channel string `json:"channel" binding:"required"`
	ClientURL  string `json:"clientUrl" binding:"required"`
}

type PublishRequest struct {
	Channel string `json:"channel" binding:"required"`
	Message  interface{} `json:"message" binding:"required"`
}
