package models

import "net/http"

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(data interface{}) *BaseResponse {
	return &BaseResponse{
		Success: true,
		Data:    data,
		Code:    http.StatusOK,
		Message: "İşlem başarılı",
	}
}

func NewEmptySuccessResponse() *BaseResponse {
	return &BaseResponse{
		Success: true,
		Data:    nil,
		Code:    http.StatusOK,
		Message: "İşlem başarılı",
	}
}

func NewErrorResponse(message string, httpCode int) *BaseResponse {
	return &BaseResponse{
		Success: false,
		Message: message,
		Code:    httpCode,
	}
}
