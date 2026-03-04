package helper

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ApiSuccess(message string, data interface{}) ApiResponse {
	return ApiResponse{Success: true, Message: message, Data: data}
}

func ApiError(message string) ApiResponse {
	return ApiResponse{Success: false, Message: message, Data: nil}
}
