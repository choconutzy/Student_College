package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerErrorResponse(err error, ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": message,
		// "error": err.Error(),
	})
}

func BadRequest(err error, ctx *gin.Context, message string, errorMsg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": message,
		"error":   errorMsg,
	})
}
