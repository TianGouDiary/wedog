package models

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"math/rand"
	"time"
)

/*
{
  "sender": "舔狗",
  "receiver": "女神",
  "content": "昨天你领完红包就把我删了，我陷入久久地沉思。我想这其中一定有什么含义，原来你是在欲擒故纵，嫌我不够爱你。无理取闹的你变得更加可爱了，我会坚守我对你的爱</mark>的。你放心好啦！今天发工资了，发了1850，给你微信转了520，支付宝1314，还剩下16。给你发了很多消息你没回。剩下16块我在小卖部买了你爱吃的老坛酸菜牛肉面，给你寄过去了。希望你保护好食欲，我去上班了爱你~~",
  "weather_info": "",
  "create_time": "2020-05-06 21:57:40"
}
*/

type WeDog struct {
	Engine      *xorm.Engine `xorm:"-"`
	dbInfo      DbInfo       `xorm:"-"`
	Id          int64        `xorm:"'id' BigInt notnull pk autoincr" json:"db_info"`
	Sender      string       `xorm:"'sender' Varchar(64)" json:"sender"`
	Receiver    string       `xorm:"'receiver' Varchar(64)" json:"receiver"`
	Content     string       `xorm:"'content' Varchar(4096)" json:"content"`
	WeatherInfo string       `xorm:"'weather_info' Varchar(16)" json:"weather_info"`
	CreateTime  time.Time    `xorm:"'create_time' DateTime" json:"create_time"`
}

type DbInfo struct {
	User     string
	Password string
	Ip       string
	Port     string
	DbName   string
}

func NewWeDog(user, password, ip, port, dbName string) (weDog *WeDog) {
	return &WeDog{
		dbInfo: DbInfo{
			User:     user,
			Password: password,
			Ip:       ip,
			Port:     port,
			DbName:   dbName,
		}}
}

func (weDog *WeDog) Open() (err error) {
	var engine *xorm.Engine
	// username:password@protocol(address)/dbname?param=value
	ping := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		weDog.dbInfo.User,
		weDog.dbInfo.Password,
		weDog.dbInfo.Ip,
		weDog.dbInfo.Port,
		weDog.dbInfo.DbName)
	if engine, err = xorm.NewEngine("mysql", ping); err != nil {
		return
	}
	fmt.Println("ping", ping)
	// 最大连接池为 20 个
	engine.SetMaxOpenConns(20)
	engine.ShowSQL(true)
	weDog.Engine = engine
	return
}

func (weDog *WeDog) Close() {
	if weDog != nil && weDog.Engine != nil {
		_ = weDog.Engine.Close()
	}
}

func (weDog *WeDog) CreateTable() (err error) {
	wd := new(WeDog)
	return weDog.Engine.CreateTables(wd)
}

func (weDog *WeDog) GetCount() (count int64, err error) {
	count, err = weDog.Engine.Count(weDog)
	if err != nil {
		return
	}
	return
}

func (weDog *WeDog) Insert(v interface{}) (err error) {
	wd, ok := v.(*WeDog)
	if !ok {
		return errors.New("type error")
	}
	_, err = weDog.Engine.Insert(wd)
	if err != nil {
		return
	}
	return
}

func (weDog *WeDog) GetRandomOne() (Model, error) {
	count, err := weDog.GetCount()
	if err != nil {
		return nil, err
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	id := r.Intn(int(count + 1))
	wd := new(WeDog)
	fmt.Println("count", count, "id", id)
	if id == 0 {
		id++
	}
	ok, err := weDog.Engine.ID(id).Get(wd)
	if !ok || err != nil {
		return nil, errors.New("Get data error")
	}
	return wd, nil
}
