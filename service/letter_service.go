package service

import (
	"math/rand"
	"wedog/common"
	"wedog/model"
	"wedog/serializer"
)

// CreateVideoService 视频投稿的服务
type LetterService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// Create 创建视频
func (service *LetterService) GetRandomLetter() serializer.Response {
	count, err := model.GetLetterCount()
	if err != common.ERR_SUCCESS {
		return serializer.Response{
			ErrCode: err.ErrorCode(),
			ErrMsg:  err.ErrorMsg(),
		}
	}
	if count == 0 {
		count = 1
	}
	id := rand.Intn(int(count))
	if id <= 0 {
		id = 1
	}
	letter, err := model.GetLetter(uint64(id))
	if err != common.ERR_SUCCESS {
		return serializer.Response{
			ErrCode: err.ErrorCode(),
			ErrMsg:  err.ErrorMsg(),
		}
	}
	return serializer.Response{
		Data: letter,
	}
}
