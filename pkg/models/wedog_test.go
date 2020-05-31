package models

import (
	"testing"
	"time"
)

const (
	User       = "pengwl"
	Password   = "123456"
	Ip         = "192.168.188.128"
	Port       = "3306"
	DbName     = "wedog"
	TimeFormat = "2006-01-02 15:04:05"
)

func TestNewWeDog(t *testing.T) {
	weDog := NewWeDog(User, Password, Ip, Port, DbName)
	t.Logf("weDog:%+v", weDog)
}

func TestWeDog_Open(t *testing.T) {
	weDog := NewWeDog(User, Password, Ip, Port, DbName)
	err := weDog.Open()
	if err != nil {
		t.Fatal(err)
	}
	weDog.Close()
}

func TestWeDog_CreateTable(t *testing.T) {
	weDog := NewWeDog(User, Password, Ip, Port, DbName)
	err := weDog.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer weDog.Close()
	err = weDog.CreateTable()
	if err != nil {
		t.Fatal(err)
	}
}

func TestWeDog_Insert(t *testing.T) {
	weDog := NewWeDog(User, Password, Ip, Port, DbName)
	err := weDog.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer weDog.Close()
	wd := WeDog{
		Sender:      "pengwl",
		Receiver:    "女神",
		Content:     "这是第4条测试",
		WeatherInfo: "晴天",
		CreateTime:  time.Now(),
	}

	err = weDog.Insert(&wd)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWeDog_GetCount(t *testing.T) {
	weDog := NewWeDog(User, Password, Ip, Port, DbName)
	err := weDog.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer weDog.Close()

	count, err := weDog.GetCount()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("count:", count)
}

func TestWeDog_GetRandomOne(t *testing.T) {
	weDog := NewWeDog(User, Password, Ip, Port, DbName)
	err := weDog.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer weDog.Close()

	wd, err := weDog.GetRandomOne()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("wd:%+v", wd)
}
