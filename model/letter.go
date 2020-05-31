package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"time"
	"wedog/common"
)

// User 用户模型
type Letter struct {
	Engine      *xorm.Engine `xorm:"-" json:"-"`
	Id          uint64       `xorm:"'id' BigInt notnull pk autoincr" json:"id"`
	Sender      string       `xorm:"'sender' Varchar(64)" json:"sender"`
	Receiver    string       `xorm:"'receiver' Varchar(64)" json:"receiver"`
	Content     string       `xorm:"'content' Varchar(4096)" json:"content"`
	WeatherInfo string       `xorm:"'weather_info' Varchar(16)" json:"weather_info"`
	CreateTime  time.Time    `xorm:"'create_time' DateTime" json:"create_time"`
}

func GetLetterCount() (int64, common.ErrConst) {
	letter := new(Letter)
	total, err := DB.Table(common.TB_LETTER_LIST).Where("id >?", 0).Count(letter)
	if err != nil {
		fmt.Print("GetLetterCount err", err)
		return 0, common.ERR_DB_FAIL
	}
	return total, common.ERR_SUCCESS
}

func GetLetter(id uint64) (*Letter, common.ErrConst) {
	letter := new(Letter)
	has, err := DB.Table(common.TB_LETTER_LIST).ID(id).Get(letter)
	if err != nil {
		fmt.Print("GetLetter err", err)
		return nil, common.ERR_DB_FAIL
	}
	if !has {
		return nil, common.ERR_DB_NO_DATA
	}
	return letter, common.ERR_SUCCESS
}