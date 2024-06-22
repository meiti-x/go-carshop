package helper

import "github.com/meiti-x/golang-web-api/api/validations"

type BaseHttpResponse struct {
	Result           any                            `json:"result"`
	Success          bool                           `json:"success"`
	ResultCode       int                            `json:"resultCode"`
	ValidationErrors *[]validations.ValidationError `json:"validationErrors,omitempty"`
	Error            any                            `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode int) BaseHttpResponse {
	return BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       resultCode,
		ValidationErrors: nil,
		Error:            nil,
	}
}
func GenerateBaseResponseWithError(result any, success bool, resultCode int, err error) BaseHttpResponse {
	return BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       resultCode,
		ValidationErrors: nil,
		Error:            err.Error(),
	}
}

func GenerateBaseResponseWithValidtionError(result any, success bool, resultCode int, err error) BaseHttpResponse {
	return BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       resultCode,
		ValidationErrors: validations.GetValidationErrors(err),
		Error:            nil,
	}
}
