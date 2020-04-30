package kernel

import "encoding/json"

type JsonSuccess struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type JsonSuccessWithoutData struct {
	Success bool `json:"success"`
}

type JsonError struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JsonErrorWithoutData struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

func JsonSuccessEncode(data interface{}) ([]byte, error) {
	if data == nil {
		return json.Marshal(&JsonSuccessWithoutData{
			Success: true,
		})
	} else {
		return json.Marshal(&JsonSuccess{
			Success: true,
			Data:    data,
		})
	}
}

func JsonErrorEncode(code uint, message string, data interface{}) ([]byte, error) {
	if data == nil {
		return json.Marshal(&JsonErrorWithoutData{
			Code:    code,
			Message: message,
		})
	} else {
		return json.Marshal(&JsonError{
			Code:    code,
			Message: message,
			Data:    data,
		})
	}
}
