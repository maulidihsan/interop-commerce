package models

type Response struct {
	Success bool `json:"success"`
	Code int32 `json:"code"`
	Message string `json:"message"`
}