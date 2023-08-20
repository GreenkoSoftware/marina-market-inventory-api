package common_function

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	constants "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/constant"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/models"
)

func CreateResponse(c *gin.Context, data interface{}, err error) models.Response {
	var response models.Response

	if err != nil {
		response = createErrorResponse(err)
	} else {
		if !isEmptyValue(data) {
			response = createSuccessResponse(data)
		} else {
			response = createNoContentResponse()
		}
	}

	SendResponse(c, response)
	return response
}

func createErrorResponse(err error) models.Response {
	status := "Error"
	code := http.StatusBadRequest
	messages := err.Error()

	if err == constants.ErrorInPassword {
		status = "Unauthorized"
		code = http.StatusUnauthorized
	}

	return models.Response{
		Status:   status,
		Code:     code,
		Messages: messages,
		Data:     nil,
	}
}

func createSuccessResponse(data interface{}) models.Response {
	return models.Response{
		Status:   "OK",
		Code:     http.StatusOK,
		Messages: "",
		Data:     &data,
	}
}

func createNoContentResponse() models.Response {
	return models.Response{
		Status:   "No Content",
		Code:     http.StatusNoContent,
		Messages: "",
		Data:     nil,
	}
}

func isEmptyValue(value interface{}) bool {
	return reflect.ValueOf(value).IsNil() || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil())
}

func SendResponse(c *gin.Context, response models.Response) {
	c.JSON(response.Code, response)
}
