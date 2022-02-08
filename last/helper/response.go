package helper

import "strings"

// REsponse used for static shaped json return
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// Empty object is used hen data ant to be null on json
type EmptyObj struct{}

//BuildRespons method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Success: status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

//BuildRespons method is to inject data value to dynamic failed response
func BuildErorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Success: false,
		Message: message,
		Error:   splittedError,
		Data:    data,
	}
	return res
}
