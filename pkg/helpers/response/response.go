package response

type SuccessResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
}

func NewSuccessResponse(code int, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Code:    code,
		Data:    data,
	}
}

func NewErrorResponse(code int, data interface{}) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Code:    code,
		Errors:  data,
	}
}
