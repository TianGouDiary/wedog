package models

type Model interface {
	Open() error
	Close()
	CreateTable() error
	GetCount() (int64, error)
	Insert(v interface{}) error
	GetRandomOne() (Model, error)
}
