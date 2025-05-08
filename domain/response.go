package domain

type ErrorResponse struct {
	Code		int				`json:"code"`
	Message	string		`json:"message,omitempty"`
	Detail	any				`json:"detail,omitempty"`
}

type SuccessResponse struct {
	Code		int				`json:"code"`
	Message	string		`json:"message,omitempty"`
	Data		any				`json:"data"`
}