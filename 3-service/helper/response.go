package helper

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
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

// 响应失败
func BuildErrorResponse(code int, message string, error interface{}, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
		Error:   error,
	}
}
