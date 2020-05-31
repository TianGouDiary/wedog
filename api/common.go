package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wedog/common"
)

type Response struct {
	ErrCode uint32
	ErrMsg  string
	Data    interface{}
}

func Return(c *gin.Context, err common.ErrConst, data interface{}) {
	if err != common.ERR_SUCCESS {
		c.JSON(http.StatusOK, Response{
			ErrCode: err.Code,
			ErrMsg:  err.Msg,
			Data:    "",
		})
	} else {
		c.JSON(http.StatusOK, Response{
			ErrCode: err.Code,
			ErrMsg:  err.Msg,
			Data:    data,
		})
	}
	return
}
