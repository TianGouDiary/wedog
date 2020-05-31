package api

import (
	"github.com/gin-gonic/gin"
	"wedog/service"
)

func GetRandomLetter(c *gin.Context) {
	letterService := service.LetterService{}
	res := letterService.GetRandomLetter()
	c.JSON(200, res)
}
