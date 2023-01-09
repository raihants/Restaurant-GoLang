package function

import (
	"encoding/json"
	"net/http"
)

// BaseRespons base response
type BaseResponse struct {
	StatusCode int         `json:"status_code" example:"200"`
	Message    string      `json:"message" example:"success"`
	Data       interface{} `json:"data"`
}

type BaseResponseErrDefault struct {
	StatusCode int      `json:"status_code" example:"500"`
	Message    string   `json:"message" example:"something went wrong, please try again"`
	Data       struct{} `json:"data"`
}

type BaseResponseErr400 struct {
	StatusCode int      `json:"status_code" example:"400"`
	Message    string   `json:"message" example:"refer to error validation"`
	Data       struct{} `json:"data"`
}

type BaseResponseErr500 struct {
	StatusCode int      `json:"status_code" example:"500"`
	Message    string   `json:"message" example:"internal server error"`
	Data       struct{} `json:"data"`
}

// SendResponse ..
func SendResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) string {
	res := new(BaseResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res.StatusCode = statusCode
	res.Message = message
	res.Data = data
	json.NewEncoder(w).Encode(res)
	return message
}
