package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error":   true,
		"message": message,
	})
}

func WriteToResponseBody(c *gin.Context, response interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

type WebResponse struct {
	Error   bool        `json:"error,omitempty"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
