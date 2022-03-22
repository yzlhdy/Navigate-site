package helper

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}
type ResponsePage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Total   int64       `json:"total"`
}

// 空对象响应
type EmptyObjectResponse struct{}

// 响应成功

func BuildResponse(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
		Error:   nil,
	}
}

func BuildResponsePage(code int, message string, data interface{}, total int64) ResponsePage {
	return ResponsePage{
		Code:    code,
		Message: message,
		Data:    data,
		Error:   nil,
		Total:   total,
	}
}

// 响应失败
func BuildErrorResponse(code int, message string, error interface{}, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
		Error:   error,
	}
}
