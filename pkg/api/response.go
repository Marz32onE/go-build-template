package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type NormalResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	ErrCode int    `json:"errCode"`
	Error   string `json:"error"`
}

func SendNormalResponse(c echo.Context, data interface{}) error {
	response := NormalResponse{
		Code: http.StatusOK,
		Data: data,
	}
	return c.JSON(http.StatusOK, response)
}

func SendErrorResponse(c echo.Context, errCode int, errMsg string) error {
	response := ErrorResponse{
		ErrCode: errCode,
		Error:   errMsg,
	}
	return c.JSON(errCode, response)
}
