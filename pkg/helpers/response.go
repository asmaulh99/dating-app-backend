package helpers

import (
	"net/http"

	"github.com/asmaulh99/dating-app-backend/pkg/configs"
	"github.com/asmaulh99/dating-app-backend/pkg/errors"
	"github.com/gin-gonic/gin"
)

// DefaultResponse represent the response success struct
type DefaultResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// ResponseError represent the response error struct
type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func JSON(ctx *gin.Context, statusCode int, data interface{}) {
	response := DefaultResponse{
		Status: http.StatusText(statusCode),
		Data:   data,
	}

	ctx.JSON(statusCode, response)
}

func SendJSONErrorResponse(ctx *gin.Context, err error) {
	var statusCode = 0
	var httpStatus string
	var message string

	cfg := configs.GetConfig()

	if dynamicErr, ok := err.(*errors.DynamicError); ok {
		statusCode = dynamicErr.StatusCode
		httpStatus = http.StatusText(statusCode)
	}

	if err == nil || cfg.Env == "production" {
		message = "Internal Server Error"
	} else {
		message = err.Error()
	}

	if statusCode == 0 {
		httpStatus = http.StatusText(500)
		statusCode = 500
	}

	response := ResponseError{
		Status:  httpStatus,
		Message: message,
	}

	ctx.AbortWithStatusJSON(statusCode, response)
}
