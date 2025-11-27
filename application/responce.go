package application

import (
	"github.com/MiladCodeLab/simple-api/dto"
	"github.com/gin-gonic/gin"
)

func JSONError(c *gin.Context, status int, message string) {
	c.JSON(status, dto.ErrorResponse{
		Error: message,
	})
}

func JSONSuccess(c *gin.Context, status int, payload any) {
	c.JSON(status, payload)
}
