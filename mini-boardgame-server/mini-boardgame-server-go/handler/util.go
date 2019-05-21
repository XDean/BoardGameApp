package handler

type J map[string]interface{}

type MessageResponse struct {
	message string
}

func M(msg string) MessageResponse {
	return MessageResponse{
		message: msg,
	}
}
