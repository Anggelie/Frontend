package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginSuccessData struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: "OK",
		Data:    data,
	}
}

func NewErrorResponse(msg string) APIResponse {
	return APIResponse{
		Success: false,
		Message: msg,
	}
}
