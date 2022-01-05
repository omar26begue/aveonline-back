package models

type HTTPResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
